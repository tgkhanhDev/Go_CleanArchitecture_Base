package api

import (
	"AuthService/internal/api/controller"
	"AuthService/internal/api/router"
	"AuthService/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Router struct {
	engine         *gin.Engine
	authController *controller.AuthController
}

func NewRouter(
	authCtl *controller.AuthController,
) *Router {
	r := &Router{
		engine:         gin.Default(),
		authController: authCtl,
	}

	// Thêm middleware CORS
	r.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//Setup static file serving
	r.Engine().Static("/swagger", "./docs/swagger-ui")
	//r.Engine().Static("/static", "./uploads")
	r.RegisterRoutes()
	logger.GetLogger().Debug("Successfully initialized Router")
	return r
}

func (r *Router) RegisterRoutes() {
	apiContext := r.engine.Group("/api/v1")
	router.RegisterAuthRoute(apiContext, r.authController)
}

// Serve khởi động HTTP server.
func (r *Router) Serve(addr string) error {
	return r.engine.Run(addr)
}

func (r *Router) Engine() *gin.Engine {
	return r.engine
}
