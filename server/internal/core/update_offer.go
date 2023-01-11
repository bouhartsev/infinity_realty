package core

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) UpdateOffer(ctx context.Context, req *domain.UpdateOfferRequest) (*domain.GetOfferResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	_, err := c.GetOffer(ctx, req.OfferId)
	if err != nil {
		return nil, err
	}

	err = c.db.UpdateOffer(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetOffer(ctx, req.OfferId)
}
