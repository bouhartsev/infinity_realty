package database

import "context"

func (d *Database) DeleteUser(ctx context.Context, id int) error {
	_, err := d.Conn.Exec(ctx, "delete from users where id = $1", id)
	return err
}
