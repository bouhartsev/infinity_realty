package database

import "context"

func (d *Database) DeleteProperty(ctx context.Context, id int) error {
	_, err := d.Conn.ExecContext(ctx, "delete from properties where id = ?", id)
	return err
}
