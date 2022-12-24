package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) GetUser(ctx context.Context, id int) (*domain.GetUserResponse, error) {
	user, err := c.db.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.UserNotFoundError
		}
		return nil, errdomain.NewInternalError(err.Error())
	}
	return &domain.GetUserResponse{User: user}, nil
}
