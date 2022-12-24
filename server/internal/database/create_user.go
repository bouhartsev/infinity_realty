package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) CreateUser(ctx context.Context, req *domain.CreateUserRequest) (int, error) {
	var createdId int
	err := d.Conn.QueryRow(ctx, `insert into users(role, name, surname, patronymic, tel, email, commission, password)
							 values($1, $2, $3, $4, $5, $6, $7, $8)
							 returning id`,
		req.Role,
		req.Name,
		req.Surname,
		req.Patronymic,
		req.Telephone,
		req.Email,
		req.Commission,
		req.Password,
	).Scan(&createdId)
	if err != nil {
		return 0, err
	}

	return createdId, nil
}
