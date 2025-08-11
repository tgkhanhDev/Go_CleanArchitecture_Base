package controllers

import (
	request "gin/internal/application/dtos/request"
	service "gin/internal/infrastructure/services"
	apiRes "gin/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: &authService}
}

func (ac *AuthController) LoginHandler(ctx *gin.Context) {

	var req request.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, apiRes.BadRequestResponse(err.Error()))
		return
	}

	// gọi services xử lý
	result, err := ac.authService.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apiRes.BadRequestResponse(err.Error()))
		return
	}

	// trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": result})
}
