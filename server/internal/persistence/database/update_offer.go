package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) UpdateOffer(ctx context.Context, req *domain.UpdateOfferRequest) error {
	var (
		args   = make([]any, 0)
		values []string
	)

	if req.ClientId != nil {
		args = append(args, *req.ClientId)
		values = append(values, "client_id = ?")
	}

	if req.RealtorId != nil {
		args = append(args, *req.RealtorId)
		values = append(values, "realtor_id = ?")
	}

	if req.PropertyId != nil {
		args = append(args, *req.PropertyId)
		values = append(values, "property_id = ?")
	}

	if req.Price != nil {
		args = append(args, *req.Price)
		values = append(values, "price = ?")
	}

	q := strings.Join(values, ",")
	query := fmt.Sprintf("update offers set %s where id = ?", q)
	args = append(args, req.OfferId)

	_, err := d.Conn.ExecContext(ctx, query, args...)
	return err
}
