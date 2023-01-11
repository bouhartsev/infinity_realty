package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (d *Database) CheckClientDependencies(ctx context.Context, id int) error {
	var userId int
	err := d.Conn.QueryRowContext(ctx, "select id from needs where client_id = ?", id).Scan(&userId)
	if err == nil {
		return ErrClientHasNeeds
	}
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	return nil
}
