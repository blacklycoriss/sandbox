package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitPostgres(url string) {
	Pool, _ = pgxpool.New(context.Background(), url) // Add error handling
}
