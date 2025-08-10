package router

import (
	"fmt"
	"gin/internal/controller"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthController *controller.AuthController
	//userCtl controller.UserController
}

func SetupRouter(cfg RouterConfig) *gin.Engine {
	r := gin.Default()
	auth := cfg.AuthController
	//user := cfg.userCtl

	// auth routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", auth.LoginHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"data":    nil,
			"message": "Route not found",
			"code":    404,
		})
	})

	fmt.Println("Nothing will goes here")
	return r

}
