package database

import "context"

func (d *Database) CheckClientTelephone(ctx context.Context, telephone string) error {
	var userId int
	row := d.Conn.QueryRowContext(ctx, "select id from clients where lower(tel) = lower(?)", telephone)
	err := row.Scan(&userId)
	return err
}
