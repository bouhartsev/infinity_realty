package core

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) CreateOffer(ctx context.Context, req *domain.CreateOfferRequest) (*domain.GetOfferResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	id, err := c.db.CreateOffer(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetOffer(ctx, id)
}
