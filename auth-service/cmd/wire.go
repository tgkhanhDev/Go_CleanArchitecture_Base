//go:build wireinject
// +build wireinject

package main

import (
	"AuthService/config"
	"AuthService/internal/api"
	"AuthService/internal/api/controller"
	"AuthService/internal/domain/entities"
	"AuthService/internal/infrastructure/persistence/databases"
	"AuthService/internal/infrastructure/persistence/repositories"
	"AuthService/internal/infrastructure/service"
	"AuthService/pkg/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Server struct holds all initialized dependencies
type Server struct {
	Router *api.Router
	DB     *gorm.DB
	Logger *zap.Logger
}

// ProvideLogger initializes the logger
func ProvideLogger() *zap.Logger {
	return logger.GetLogger()
}

// ProvideConfig loads the environment config
func ProvideConfig() (*config.EnvProps, error) {
	return config.LoadEnvConfig()
}

// ProvideDB initializes the database connection
func ProvideDB(cfg *config.EnvProps) (*gorm.DB, error) {
	pgdb, err := databases.NewDatabasePgConnection(cfg.DatabaseProps)
	if err != nil {
		logger.GetLogger().Error("Failed when initiating connection to PostgreSQL: " + err.Error())
		return nil, err
	}
	return pgdb, nil
}

// NewServer constructs the Server
func NewServer(router *api.Router, db *gorm.DB, log *zap.Logger) *Server {
	return &Server{
		Router: router,
		DB:     db,
		Logger: log,
	}
}

// GENERIC PROVIDER
func ProvideGenericAccountRepository(db *gorm.DB) repositories.GenericRepository[entities.Account] {
	return repositories.NewGenericRepository[entities.Account](db)
}

// === Provider Sets for Organization ===

// InfrastructureSet cung cấp các implementation cụ thể
var InfrastructureSet = wire.NewSet(
	ProvideConfig,
	ProvideLogger,
	ProvideDB,
	ProvideGenericAccountRepository,
	repositories.NewUserRepository,
)

// ApplicationSet cung cấp các service/use-case của tầng application
var ApplicationSet = wire.NewSet(
	service.NewAuthService,
	service.NewUserService,
)

// APISet cung cấp các controller và router
var APISet = wire.NewSet(
	controller.NewAuthController,
	controller.NewUserController,
	api.NewRouter,
)

// InitializeServer is the Wire injector
func InitializeServer() (*Server, error) {
	wire.Build(
		InfrastructureSet,
		ApplicationSet,
		APISet,
		NewServer,
	)
	return nil, nil
}
