package controller

import (
	"AuthService/internal/application/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: &userService,
	}
}

func (uc *UserController) GetAllUserHandler(context *gin.Context) {
	context.JSON(200, nil)
}
