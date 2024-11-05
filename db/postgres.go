package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

var DB *sql.DB

// Connect initialises a connection to the PostgreSQL database
func Connect() error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s??sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Connected to PostgreSQL successfully")
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
