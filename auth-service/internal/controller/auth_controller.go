package controller

import (
	dto2 "gin/internal/dto"
	dto "gin/internal/dto/request"
	service "gin/internal/service/impl"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: &authService}
}

func (ac *AuthController) LoginHandler(ctx *gin.Context) {

	var req dto.LoginRequest

	// Check body rỗng
	bodyBytes, _ := io.ReadAll(ctx.Request.Body)
	if len(bodyBytes) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Request body is empty",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto2.BadRequestResponse(err.Error()))
		return
	}

	// gọi service xử lý
	result, err := ac.authService.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto2.BadRequestResponse(err.Error()))
		return
	}

	// trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": result})
}
