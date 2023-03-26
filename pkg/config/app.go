package config

import (
	"github.com/spf13/viper"
)

// HINT: holds on general app settings
type AppConfig struct {
	Port    string `dotenv:"APP_PORT"`
	Mode    string `dotenv:"APP_MODE"`
	GoEnv   string `dotenv:"GO_ENV"`
	Version string `dotenv:"VERSION"`
}

func NewAppConfig() *AppConfig {
	var config AppConfig

	// Setting defaults
	if viper.Get("GO_ENV") == nil {
		viper.SetDefault("GO_ENV", "local")
	}

	if viper.Get("VERSION") == nil {
		viper.SetDefault("VERSION", "local")
	}

	viper.Unmarshal(&config)

	return &config
}

func (ac AppConfig) IsProduction() bool {
	if ac.Mode == "develop" {
		return false
	}

	return true
}
