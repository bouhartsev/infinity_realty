package database

import (
	"context"
)

func (d *Database) CheckClientEmail(ctx context.Context, email string) error {
	var userId int
	row := d.Conn.QueryRow(ctx, "select id from clients where lower(email) = lower($1)", email)
	err := row.Scan(&userId)
	return err
}
