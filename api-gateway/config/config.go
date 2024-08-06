package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AuthServiceHost string
	AuthServicePort string

	MedalServiceHost string
	MedalServicePort string

	ApiGatewayHost string
	ApiGatewayPort string
}

func Load() Config {
	godotenv.Load()
	cfg := viper.New()
	cfg.AutomaticEnv()

	conf := Config{
		AuthServiceHost: cfg.GetString("AUTH_SERVICE_HOST"),
		AuthServicePort: cfg.GetString("AUTH_SERVICE_PORT"),

		MedalServiceHost: cfg.GetString("MEDAL_SERVICE_HOST"),
		MedalServicePort: cfg.GetString("MEDAL_SERVICE_PORT"),

		ApiGatewayHost: cfg.GetString("API_GATEWAY_HOST"),
		ApiGatewayPort: cfg.GetString("API_GATEWAY_PORT"),
	}

	return conf
}
