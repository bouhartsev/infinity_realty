package database

import "context"

func (d *Database) DeleteClient(ctx context.Context, id int) error {
	_, err := d.Conn.ExecContext(ctx, "delete from clients where id = ?", id)
	return err
}
