package postgres

import (
	"database/sql"
	"fmt"
	"user-service/config"

	_ "github.com/lib/pq"
)

func InitDB(cfg config.Config) (*sql.DB, error) {
	target := fmt.Sprintf(
		`
			host=%s
			user=%s
			password=%s
			dbname=%s
			sslmode=disable	
		`,
		cfg.Postgres.PostgresHost,
		cfg.Postgres.PostgresUser,
		cfg.Postgres.PostgresPassword,
		cfg.Postgres.PostgresDatabase,
	)

	db, err := sql.Open("postgres", target)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
