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

type Redis struct {
	RedisHost string
	RedisPort string
}

type Config struct {
	Postgres Postgres

	Redis Redis

	UserServiceHost string
	UserServicePort string
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

		Redis: Redis{
			RedisHost: cfg.GetString("REDIS_HOST"),
			RedisPort: cfg.GetString("REDIS_PORT"),
		},

		UserServiceHost: cfg.GetString("USER_SERVICE_HOST"),
		UserServicePort: cfg.GetString("USER_SERVICE_PORT"),
	}

	return conf
}
