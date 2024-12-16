package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/go-redis/redis/v8"
)

type Database struct {
	Postgres *sql.DB
	Redis    *redis.Client
}

func NewDatabase() (*Database, error) {
	// PostgreSQL connection
	pgConn := "host=localhost port=5432 user=your_user dbname=your_db password=your_password sslmode=disable"
	pg, err := sql.Open("postgres", pgConn)
	if err != nil {
		return nil, err
	}

	// Redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Database{
		Postgres: pg,
		Redis:    rdb,
	}, nil
} 