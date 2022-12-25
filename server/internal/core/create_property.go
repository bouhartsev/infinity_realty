package core

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) CreateProperty(ctx context.Context, req *domain.CreatePropertyRequest) (*domain.GetPropertyResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := c.db.CreateProperty(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetProperty(ctx, id)
}
