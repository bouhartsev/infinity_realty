package database

import "context"

func (d *Database) DeleteRealtor(ctx context.Context, id int) error {
	_, err := d.Conn.ExecContext(ctx, "delete from realtors where id = ?", id)
	return err
}
