package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) GetUser(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	err := d.Conn.QueryRow(ctx, `select id, role, name, surname, patronymic, tel, email, commission
                              from users
                              where id = $1`,
		id).Scan(
		&user.Id,
		&user.Role,
		&user.Name,
		&user.Surname,
		&user.Patronymic,
		&user.Telephone,
		&user.Email,
		&user.Commission,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
