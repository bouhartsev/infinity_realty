package database

import "context"

func (d *Database) DeleteProperty(ctx context.Context, id int) error {
	_, err := d.Conn.Exec(ctx, "delete from properties where id = $1", id)
	return err
}
