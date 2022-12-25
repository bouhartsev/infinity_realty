package database

import "context"

func (d *Database) CheckClientTelephone(ctx context.Context, telephone string) error {
	var userId int
	row := d.Conn.QueryRow(ctx, "select id from clients where lower(tel) = lower($1)", telephone)
	err := row.Scan(&userId)
	return err
}
