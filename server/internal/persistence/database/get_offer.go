package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) GetOffer(ctx context.Context, id int) (*domain.Offer, error) {
	var resp domain.Offer
	err := d.Conn.QueryRowContext(ctx, `select id, client_id, realtor_id, property_id, price from offers where id = ?`, id).Scan(
		&resp.Id,
		&resp.ClientId,
		&resp.RealtorId,
		&resp.PropertyId,
		&resp.Price,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
