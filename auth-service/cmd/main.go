package main

import (
	"database/sql"
	"gin/internal/config"
	"gin/internal/controller"
	persistence "gin/internal/repository/impl"
	"gin/internal/router"
	service "gin/internal/service/impl"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	cfg := config.LoadDatabaseConfig()

	db, err := sql.Open("postgres", cfg.GetPostgresDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Manual Dependency Injection
	userRepo := persistence.NewUserRepository(db)
	// 3. Service
	authService := service.NewAuthService(userRepo)
	// 4. Controller
	authController := controller.NewAuthController(*authService)

	// 5. Route
	r := router.SetupRouter(router.RouterConfig{
		AuthController: authController,
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
