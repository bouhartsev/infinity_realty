package database

import "context"

func (d *Database) DeleteOffer(ctx context.Context, id int) error {
	_, err := d.Conn.ExecContext(ctx, "delete from offers where id = ?", id)
	return err
}
