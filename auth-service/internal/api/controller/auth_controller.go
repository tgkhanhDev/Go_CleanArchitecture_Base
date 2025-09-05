package controller

import (
	"AuthService/internal/application/dtos/request"
	"AuthService/internal/application/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(
	authService services.AuthService,
) *AuthController {
	return &AuthController{
		authService: &authService,
	}
}

func (ac *AuthController) GetAllAccountsHandler(context *gin.Context) {
	accounts, err := (*ac.authService).GetAllAccounts()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, accounts)
}

func (ac *AuthController) LoginHandler(context *gin.Context) {
	var loginReq request.LoginRequest
	if err := context.ShouldBindJSON(&loginReq); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs := make(map[string]string)
			for _, fe := range ve {
				switch fe.Tag() {
				case "required":
					errs[fe.Field()] = fe.Field() + " is required"
				case "email":
					errs[fe.Field()] = "Invalid email format"
				case "min":
					errs[fe.Field()] = fe.Field() + " must be at least " + fe.Param() + " characters"
				default:
					errs[fe.Field()] = fe.Error()
				}
			}
			context.JSON(400, gin.H{"data": nil, "message": errs, "code": 400})
			return
		}
		context.JSON(400, gin.H{"data": nil, "message": err.Error(), "code": 400})
		return
	}
	loginResp, err := (*ac.authService).Login(loginReq)
	if err != nil {
		context.JSON(401, gin.H{"data": nil, "message": err.Error(), "code": 401})
		return
	}

	refreshToken := loginResp.RefreshToken
	context.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/", "", false, true)
	context.JSON(200, gin.H{"data": loginResp, "message": "Login successful", "code": 200})
}

func (ac *AuthController) RegisterHandler(context *gin.Context) {
	// Implementation for user registration
	var createAccountReq request.CreateAccountRequest
	if err := context.ShouldBindJSON(&createAccountReq); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs := make(map[string]string)
			for _, fe := range ve {
				switch fe.Tag() {
				case "required":
					errs[fe.Field()] = fe.Field() + " is required"
				case "email":
					errs[fe.Field()] = "Invalid email format"
				case "min":
					errs[fe.Field()] = fe.Field() + " must be at least " + fe.Param() + " characters"
				default:
					errs[fe.Field()] = fe.Error()
				}
			}
			context.JSON(400, gin.H{"data": nil, "message": errs, "code": 400})
			return
		}
		context.JSON(400, gin.H{"data": nil, "message": err.Error(), "code": 400})
		return
	}
	registerResp, err := (*ac.authService).CreateAccount(createAccountReq)
	if err != nil {
		context.JSON(500, gin.H{"data": nil, "message": err.Error(), "code": 500})
		return
	}
	context.JSON(201, gin.H{"data": registerResp, "message": "Account created successfully", "code": 201})
}

func (ac *AuthController) LogoutHandler(context *gin.Context) {
	context.SetCookie("refresh_token", "", -1, "/", "", false, true)
	context.JSON(200, gin.H{"data": nil, "message": "Logout successful", "code": 200})
}
