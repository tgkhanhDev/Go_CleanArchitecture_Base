package service

import (
	"AuthService/internal/application/dtos/request"
	"AuthService/internal/application/dtos/response"
	"AuthService/internal/application/interfaces"
	"AuthService/internal/application/services"
	"AuthService/internal/domain/entities"
	"AuthService/pkg/logger"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
)

type authServiceImpl struct {
	accountRepo interfaces.GenericRepository[entities.Account]
	db          *gorm.DB
}

func NewAuthService(db *gorm.DB, accountRepo interfaces.GenericRepository[entities.Account]) services.AuthService {
	return &authServiceImpl{
		db:          db,
		accountRepo: accountRepo,
	}
}

func (a *authServiceImpl) GetAllAccounts() ([]entities.Account, error) {
	ctx := context.Background()
	accounts, _, err := a.accountRepo.GetAll(ctx, nil, nil, 0, 0)
	if err != nil {
		log := logger.GetLogger()
		log.Error("Error fetching accounts: " + err.Error())
		return nil, err
	}
	return accounts, nil
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func generateToken(account *entities.Account, duration time.Duration) (string, int64, error) {
	expiresAt := time.Now().Add(duration).Unix()
	claims := jwt.MapClaims{
		"sub":   account.ID.String(),
		"email": account.Email,
		"exp":   expiresAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", 0, err
	}
	return signedToken, expiresAt, nil
}

func (a *authServiceImpl) Login(request request.LoginRequest) (*response.LoginResponse, error) {
	ctx := context.Background()
	filter := func(db *gorm.DB) *gorm.DB {
		return db.Where("email = ?", request.Email)
	}
	account, err := a.accountRepo.GetByCondition(ctx, filter, nil, false)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(request.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	accessToken, expiresAt, err := generateToken(account, 15*time.Minute)
	if err != nil {
		return nil, err
	}
	refreshToken, _, err := generateToken(account, 7*24*time.Hour)
	if err != nil {
		return nil, err
	}
	resp := &response.LoginResponse{
		AccessToken:  accessToken,
		ExpiresAt:    expiresAt,
		RefreshToken: refreshToken,
	}
	return resp, nil
}

func (a *authServiceImpl) CreateAccount(req request.CreateAccountRequest) (*entities.Account, error) {
	// Hash the password before saving
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	accountModel := &entities.Account{
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         "CUSTOMER",
		IsActive:     true,
	}

	if err := a.db.Create(accountModel).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
