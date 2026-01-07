package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var Pool *pgxpool.Pool

func InitializeDatabase() {
	connectionString := os.Getenv("POSTGRES_CONN_STRING")

	var err error
	Pool, err = pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}

}

func RunMigrations(migrationsDir string) {
	// Use stdlib to get a *sql.DB from the pool's config
	// This is compatible with goose
	db := stdlib.OpenDBFromPool(Pool)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("Could not set goose dialect: ", err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatal("Could not run migrations: ", err)
	}

	log.Println("Migrations completed successfully")
}

func Close() {
	if Pool != nil {
		Pool.Close()
	}
}