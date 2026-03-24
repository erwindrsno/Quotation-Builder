package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // The underscore is required here
)

func InitDB() *sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSLMODE")

	datasource := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, ssl,
	)

	db, err := sql.Open("postgres", datasource)
	if err != nil {
		log.Fatal(err)
	}

	// Maximum number of open connections to the database.
	db.SetMaxOpenConns(25)

	// Maximum number of idle connections retained in the pool.
	db.SetMaxIdleConns(10)

	// Maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(5 * time.Minute)

	// Maximum amount of time a connection may sit idle before being closed.
	db.SetConnMaxIdleTime(1 * time.Minute)

	// Verify the connection is alive
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("DB is up.\n")
	}
	return db
}
