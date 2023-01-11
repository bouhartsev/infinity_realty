package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (d *Database) CheckPropertyDependency(ctx context.Context, id int) error {
	var propId int
	err := d.Conn.QueryRowContext(ctx, "select id from offers where property_id = ?", id).Scan(&propId)
	if err == nil {
		return ErrPropertyHasOffers
	}
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	return nil
}
