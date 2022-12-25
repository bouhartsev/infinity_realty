package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) CreateRealtor(ctx context.Context, req *domain.CreateRealtorRequest) (int, error) {
	var createdId int
	err := d.Conn.QueryRow(ctx, `insert into realtors(name, surname, patronymic, commission)
							 values($1, $2, $3, $4)
							 returning id`,
		req.Name,
		req.Surname,
		req.Patronymic,
		req.Commission,
	).Scan(&createdId)
	if err != nil {
		return 0, err
	}

	return createdId, nil
}
