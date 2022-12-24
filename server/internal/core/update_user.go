package core

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) UpdateUser(ctx context.Context, req *domain.UpdateUserRequest) (*domain.GetUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	_, err := c.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	err = c.db.UpdateUser(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetUser(ctx, req.UserId)
}
