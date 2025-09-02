package databases

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseProps struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func NewDatabasePgConnection(cfg DatabaseProps) (*gorm.DB, error) {
	props := cfg
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", props.DBUser, props.DBPassword, props.DBHost, props.DBPort, props.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}
	return db, nil
}
