package router

import (
	"AuthService/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoute(parentGrp *gin.RouterGroup, ctl *controller.AuthController) {
	authGroup := parentGrp.Group("/auth")
	{
		authGroup.GET("/account", ctl.GetAllAccountsHandler)
		//parentGrp.POST("/login", ctl.)
	}
}
