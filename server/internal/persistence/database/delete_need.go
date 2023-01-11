package database

import "context"

func (d *Database) DeleteNeed(ctx context.Context, id int) error {
	_, err := d.Conn.ExecContext(ctx, "delete from needs where id = ?", id)
	return err
}
