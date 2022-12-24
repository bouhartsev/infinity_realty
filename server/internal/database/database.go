package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(connString string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
