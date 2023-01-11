package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) CreateClient(ctx context.Context, req *domain.CreateClientRequest) (int, error) {
	res, err := d.Conn.ExecContext(ctx, `insert into clients(name, surname, patronymic, tel, email)
							 values(?, ?, ?, ?, ?)`,
		req.Name,
		req.Surname,
		req.Patronymic,
		req.Telephone,
		req.Email,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
