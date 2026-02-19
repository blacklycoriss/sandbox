package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitPostgres(url string) *pgxpool.Pool {
	Pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		panic("Failed to connect to Postgres: " + err.Error())
	}
	return Pool
}
