package config

import (
	"AuthService/internal/infrastructure/persistence/databases"
	"AuthService/pkg/logger"
	"github.com/spf13/viper"
)

type EnvProps struct {
	PublicFilePath  string `mapstructure:"PUBLIC_FILE_PATH"`
	PrivateFilePath string `mapstructure:"PRIVATE_FILE_PATH"`
	AppPort         string `mapstructure:"APP_PORT"`
	JwtSecretKey    string `mapstructure:"JWT_SECRET"`
	//DB Configs
	databases.DatabaseProps `mapstructure:",squash"`
}

var envConfig *EnvProps

func LoadEnvConfig() (*EnvProps, error) {
	viper.SetConfigFile("app.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logger.GetLogger().Error("Failed to read config file: " + err.Error())
		return nil, err
	}

	var config EnvProps
	// Unmarshal file paths
	if err := viper.Sub("env_vars").Unmarshal(&config); err != nil {
		logger.GetLogger().Error("Failed to unmarshal local environments: " + err.Error())
		return nil, err
	}
	// Unmarshal database config
	if err := viper.Sub("env_vars.database_config").Unmarshal(&config); err != nil {
		logger.GetLogger().Error("Failed to unmarshal database config: " + err.Error())
		return nil, err
	}

	envConfig = &config
	return envConfig, nil
}

func GetConfig() *EnvProps {
	return envConfig
}
