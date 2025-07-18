package store

import (
	"fmt"
	"log"
	"mocku/config"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func NewConnection(cfg config.EnvVar) (Queryable, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.PGHost, cfg.PGPort, cfg.PGUser, cfg.PGPassword, cfg.PGDatabase, "disable")
	connection, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Println("Failed to open connection: %v", err)
		return nil, err
	}

	if err := connection.Ping(); err != nil {
		log.Println("Failed to ping connection: %v", err)
		return nil, err
	}

	return connection, nil
}
