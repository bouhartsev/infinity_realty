package database

import "github.com/jackc/pgx"

func NewConnection(connString string) (*pgx.ConnPool, error) {
	cfg, err := pgx.ParseConnectionString(connString)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     cfg,
		MaxConnections: 25,
	})
	if err != nil {
		return nil, err
	}

	return conn, nil
}
