package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) CreateClient(ctx context.Context, req *domain.CreateClientRequest) (*domain.GetClientResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.Email != nil {
		err := c.db.CheckClientEmail(ctx, *req.Email)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.NewInternalError(err.Error())
		}
		if err == nil {
			return nil, errdomain.NewUserError("Email already taken.", "email")
		}
	}

	if req.Telephone != nil {
		err := c.db.CheckClientTelephone(ctx, *req.Telephone)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.NewInternalError(err.Error())
		}
		if err == nil {
			return nil, errdomain.NewUserError("Telephone already taken.", "telephone")
		}
	}

	id, err := c.db.CreateClient(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetClient(ctx, id)
}
