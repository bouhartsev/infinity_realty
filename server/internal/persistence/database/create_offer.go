package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) CreateOffer(ctx context.Context, req *domain.CreateOfferRequest) (int, error) {
	res, err := d.Conn.ExecContext(ctx, `insert into offers(client_id, realtor_id, property_id, price)
							 values(?, ?, ?, ?)`,
		req.ClientId,
		req.RealtorId,
		req.PropertyId,
		req.Price,
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
