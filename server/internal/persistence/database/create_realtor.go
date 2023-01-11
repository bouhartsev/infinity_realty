package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) CreateRealtor(ctx context.Context, req *domain.CreateRealtorRequest) (int, error) {
	res, err := d.Conn.ExecContext(ctx, `insert into realtors(name, surname, patronymic, commission)
							 values(?, ?, ?, ?)`,
		req.Name,
		req.Surname,
		req.Patronymic,
		req.Commission,
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
