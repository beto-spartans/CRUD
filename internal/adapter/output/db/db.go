package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error abriendo la base de datos: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error conectando con la base de datos: %w", err)
	}

	return db, nil
}
