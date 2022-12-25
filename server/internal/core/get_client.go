package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) GetClient(ctx context.Context, id int) (*domain.GetClientResponse, error) {
	user, err := c.db.GetClient(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.ClientNotFoundError
		}
		return nil, errdomain.NewInternalError(err.Error())
	}
	return &domain.GetClientResponse{Client: user}, nil
}
