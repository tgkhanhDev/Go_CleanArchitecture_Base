package controller

import (
	"AuthService/internal/application/services"
	"github.com/gin-gonic/gin"
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
