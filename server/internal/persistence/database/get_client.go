package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) GetClient(ctx context.Context, id int) (*domain.Client, error) {
	var client domain.Client
	err := d.Conn.QueryRowContext(ctx, `select id, name, surname, patronymic, tel, email from clients where id = ?`, id).Scan(
		&client.Id,
		&client.Name,
		&client.Surname,
		&client.Patronymic,
		&client.Telephone,
		&client.Email,
	)
	if err != nil {
		return nil, err
	}

	return &client, nil
}
