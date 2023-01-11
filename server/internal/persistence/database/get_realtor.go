package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) GetRealtor(ctx context.Context, id int) (*domain.Realtor, error) {
	var resp domain.Realtor
	err := d.Conn.QueryRowContext(ctx, `select id, name, surname, patronymic, commission from realtors where id = ?`, id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.Surname,
		&resp.Patronymic,
		&resp.Commission,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
