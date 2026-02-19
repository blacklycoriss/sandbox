package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitPostgres(url string) error {
	_, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return err
	}
	return nil
}
