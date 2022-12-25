package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (d *Database) CheckClientDependencies(ctx context.Context, id int) error {
	var userId int
	err := d.Conn.QueryRow(ctx, "select id from needs where client_id = $1", id).Scan(&userId)
	if err == nil {
		return ErrClientHasNeeds
	}
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	return nil
}
