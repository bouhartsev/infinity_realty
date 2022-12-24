package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) SignUp(ctx context.Context, req *domain.CreateUserRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	if req.Role == 1 {
		return errdomain.NewUserError("Incorrect role provided.", "role")
	}

	var userId int

	if req.Email != nil {
		err := c.db.QueryRow(ctx, "select id from users where lower(email) = lower($1)", req.Email).Scan(&userId)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return errdomain.NewInternalError(err.Error())
		}
		if err == nil {
			return errdomain.NewUserError("Email already taken.", "email")
		}
	}

	if req.Telephone != nil {
		err := c.db.QueryRow(ctx, "select id from users where lower(telephone) = lower($1)", req.Telephone).Scan(&userId)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return errdomain.NewInternalError(err.Error())
		}
		if err == nil {
			return errdomain.NewUserError("Telephone already taken.", "telephone")
		}
	}

	_, err := c.db.Exec(ctx, `insert into users(role, name, surname, patronymic, tel, email, commission, password)
							 values($1, $2, $3, $4, $5, $6, $7, $8)`,
		req.Role,
		req.Name,
		req.Surname,
		req.Patronymic,
		req.Telephone,
		req.Email,
		req.Commission,
		req.Password,
	)
	if err != nil {
		return errdomain.NewInternalError(err.Error())
	}

	return nil
}
