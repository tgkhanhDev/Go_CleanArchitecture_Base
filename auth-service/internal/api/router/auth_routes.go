package router

import (
	"AuthService/internal/api/controller"
	"AuthService/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoute(parentGrp *gin.RouterGroup, ctl *controller.AuthController) {
	authGroup := parentGrp.Group("/auth")
	{
		authGroup.GET("/account", middleware.JWTAuthMiddleware(), ctl.GetAllAccountsHandler)
		authGroup.POST("/login", ctl.LoginHandler)
		authGroup.POST("/register", ctl.RegisterHandler)
		//parentGrp.POST("/logout", ctl.)
	}
}
