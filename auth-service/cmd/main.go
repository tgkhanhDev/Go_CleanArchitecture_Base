package main

import (
	"database/sql"
	"gin/internal/api"
	controller "gin/internal/api/controllers"
	persistence "gin/internal/infrastructure/persistence"
	"gin/internal/infrastructure/persistence/databases"
	service "gin/internal/infrastructure/services"
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
	r := api.NewRouter(authController)
	r.RegisterRoutes()

	// 6. Start the server
	log.Println("Starting server on port :8080")
	if err := r.Serve(":8080"); err != nil {
		log.Fatalf("could not start server: %s", err)
	}

}
