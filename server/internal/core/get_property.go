package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) GetProperty(ctx context.Context, id int) (*domain.GetPropertyResponse, error) {
	prop, err := c.db.GetProperty(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.PropertyNotFoundError
		}
		return nil, errdomain.NewInternalError(err.Error())
	}

	return &domain.GetPropertyResponse{Property: *prop}, nil
}
