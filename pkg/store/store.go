package store

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// NewConnection creates a new database connection using the provided DSN string
func NewConnection(dsn string) (Queryable, error) {
	connection, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Printf("Failed to open connection: %v", err)
		return nil, err
	}

	if err := connection.Ping(); err != nil {
		log.Printf("Failed to ping connection: %v", err)
		return nil, err
	}

	return connection, nil
}
