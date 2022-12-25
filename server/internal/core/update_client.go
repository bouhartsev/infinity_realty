package core

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) UpdateClient(ctx context.Context, req *domain.UpdateClientRequest) (*domain.GetClientResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	_, err := c.GetClient(ctx, req.ClientId)
	if err != nil {
		return nil, err
	}

	err = c.db.UpdateClient(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetClient(ctx, req.ClientId)
}
