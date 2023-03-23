package config

import (
	"github.com/spf13/viper"
)

// HINT: App configs
type Config struct {
	Port    string `dotenv:"APP_PORT"`
	GoEnv   string `dotenv:"GO_ENV"`
	Version string `dotenv:"VERSION"`
}

// NewAppConfig - creates the application config struct
func NewAppConfig() *Config {
	var config Config

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
