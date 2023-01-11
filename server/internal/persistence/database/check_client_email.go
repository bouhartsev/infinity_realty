package database

import (
	"context"
)

func (d *Database) CheckClientEmail(ctx context.Context, email string) error {
	var userId int
	row := d.Conn.QueryRowContext(ctx, "select id from clients where lower(email) = lower(?)", email)
	err := row.Scan(&userId)
	return err
}
