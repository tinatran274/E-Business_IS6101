package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port string `mapstructure:"PORT"`

	PostgresHost string `mapstructure:"POSTGRES_HOST"`

	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDBName   string `mapstructure:"POSTGRES_DB"`

	Addr     string `mapstructure:"REDIS_ADDR"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`

	JwtSecret string `mapstructure:"JWT_SECRET_KEY"`
}

func MustNewAppConfig(filePath string) *AppConfig {
	var config AppConfig
	viper.SetConfigFile(filePath)
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	return &config
}
