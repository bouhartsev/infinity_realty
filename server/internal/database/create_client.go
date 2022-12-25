package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) CreateClient(ctx context.Context, req *domain.CreateClientRequest) (int, error) {
	var createdId int
	err := d.Conn.QueryRow(ctx, `insert into clients(name, surname, patronymic, tel, email)
							 values($1, $2, $3, $4, $5)
							 returning id`,
		req.Name,
		req.Surname,
		req.Patronymic,
		req.Telephone,
		req.Email,
	).Scan(&createdId)
	if err != nil {
		return 0, err
	}

	return createdId, nil
}
