package database

import "context"

func (d *Database) DeleteClient(ctx context.Context, id int) error {
	_, err := d.Conn.Exec(ctx, "delete from clients where id = $1", id)
	return err
}
