package database

import "context"

func (d *Database) DeleteRealtor(ctx context.Context, id int) error {
	_, err := d.Conn.Exec(ctx, "delete from realtors where id = $1", id)
	return err
}
