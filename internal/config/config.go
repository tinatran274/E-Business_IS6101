package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port string `mapstructure:"PORT"`
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
