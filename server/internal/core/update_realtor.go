package core

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) UpdateRealtor(ctx context.Context, req *domain.UpdateRealtorRequest) (*domain.GetRealtorResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	_, err := c.GetRealtor(ctx, req.RealtorId)
	if err != nil {
		return nil, err
	}

	err = c.db.UpdateRealtor(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetRealtor(ctx, req.RealtorId)
}
