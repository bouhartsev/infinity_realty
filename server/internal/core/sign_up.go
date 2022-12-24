package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) SignUp(ctx context.Context, req *domain.CreateUserRequest) error {
	if req.Role == 1 {
		return errdomain.NewUserError("Incorrect role provided.", "role")
	}

	if err := req.Validate(); err != nil {
		return err
	}

	if req.Email != nil {
		err := c.db.CheckUserEmail(ctx, *req.Email)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return errdomain.NewInternalError(err.Error())
		}
		if err == nil {
			return errdomain.NewUserError("Email already taken.", "email")
		}
	}

	if req.Telephone != nil {
		err := c.db.CheckUserTelephone(ctx, *req.Telephone)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return errdomain.NewInternalError(err.Error())
		}
		if err == nil {
			return errdomain.NewUserError("Telephone already taken.", "telephone")
		}
	}

	_, err := c.db.CreateUser(ctx, req)
	if err != nil {
		return errdomain.NewInternalError(err.Error())
	}
	return err
}
