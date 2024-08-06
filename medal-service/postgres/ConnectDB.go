package postgres

import (
	"database/sql"
	"fmt"
)

func InitDB() (*sql.DB, error) {
	target := fmt.Sprintf(`
		host=%s
		user=%s
		password=%s
		dbname=%s
		sslmode=disable
	`,
		"localhost",
		"postgres",
		"1",
		"medalsdb",
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
