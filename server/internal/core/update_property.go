package core

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) UpdateProperty(ctx context.Context, req *domain.UpdatePropertyRequest) (*domain.GetPropertyResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	_, err = c.GetProperty(ctx, req.PropertyId)
	if err != nil {
		return nil, err
	}

	err = c.db.UpdateProperty(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetProperty(ctx, req.PropertyId)
}
