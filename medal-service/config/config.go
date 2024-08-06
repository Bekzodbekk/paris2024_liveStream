package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Postgres struct {
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

type Config struct {
	Postgres Postgres

	MedalServiceHost string
	MedalServicePort string
}

func Load() Config {
	godotenv.Load()
	cfg := viper.New()
	cfg.AutomaticEnv()

	conf := Config{
		Postgres: Postgres{
			PostgresHost:     cfg.GetString("POSTGRES_HOST"),
			PostgresUser:     cfg.GetString("POSTGRES_USER"),
			PostgresPassword: cfg.GetString("POSTGRES_PASSWORD"),
			PostgresDatabase: cfg.GetString("POSTGRES_DATABASE"),
		},

		MedalServiceHost: cfg.GetString("MEDAL_SERVICE_HOST"),
		MedalServicePort: cfg.GetString("MEDAL_SERVICE_PORT"),
	}

	return conf
}
