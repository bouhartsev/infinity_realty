package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.GetUserResponse, error) {
	user, err := c.getUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if user.Role != 1 {
		return nil, errdomain.ForbiddenError
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.Email != nil {
		err := c.db.CheckUserEmail(ctx, *req.Email)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.NewInternalError(err.Error())
		}
		if err == nil {
			return nil, errdomain.NewUserError("Email already taken.", "email")
		}
	}

	if req.Telephone != nil {
		err := c.db.CheckUserTelephone(ctx, *req.Telephone)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.NewInternalError(err.Error())
		}
		if err == nil {
			return nil, errdomain.NewUserError("Telephone already taken.", "telephone")
		}
	}

	id, err := c.db.CreateUser(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetUser(ctx, id)
}
