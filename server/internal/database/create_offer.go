package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) CreateOffer(ctx context.Context, req *domain.CreateOfferRequest) (int, error) {
	var createdId int
	err := d.Conn.QueryRow(ctx, `insert into offers(client_id, realtor_id, property_id, price)
							 values($1, $2, $3, $4)
							 returning id`,
		req.ClientId,
		req.RealtorId,
		req.PropertyId,
		req.Price,
	).Scan(&createdId)
	if err != nil {
		return 0, err
	}

	return createdId, nil
}
