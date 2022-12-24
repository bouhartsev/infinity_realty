package database

import (
	"context"
)

func (d *Database) CheckUserEmail(ctx context.Context, email string) error {
	var userId int
	row := d.Conn.QueryRow(ctx, "select id from users where lower(email) = lower($1)", email)
	err := row.Scan(&userId)
	return err
}
