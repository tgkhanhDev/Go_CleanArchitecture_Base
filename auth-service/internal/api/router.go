package api

import (
	"gin/internal/api/controllers"
	apiRes "gin/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	engine         *gin.Engine
	authController *controllers.AuthController
	// userController *controllers.UserController
}

// NewRouter là hàm khởi tạo cho Router.
func NewRouter(authCtl *controllers.AuthController) *Router {
	return &Router{
		engine:         gin.Default(),
		authController: authCtl,
	}
}

func (r *Router) RegisterRoutes() {
	//TODO: middlewares
	//r.engine.Use(middlewares.LoggingMiddleware())
	r.engine.Use(gin.Recovery()) // Middleware phục hồi của Gin

	// Nhóm các routes
	apiGroup := r.engine.Group("/api/v1")
	{
		authGroup := apiGroup.Group("/auth")
		{
			authGroup.POST("/login", r.authController.LoginHandler)
			// authGroup.POST("/register", r.authController.RegisterHandler)
		}

		// userGroup := apiGroup.Group("/users")
		// {
		//     userGroup.GET("/:id", r.userController.GetByID)
		// }
	}

	// Xử lý route không tồn tại (404 Not Found)
	r.engine.NoRoute(func(c *gin.Context) {
		// Sử dụng lại helper response của bạn!
		c.JSON(http.StatusNotFound, apiRes.NotFoundResponse("The requested route was not found."))
	})
}

// Serve khởi động HTTP server.
func (r *Router) Serve(addr string) error {
	return r.engine.Run(addr)
}
