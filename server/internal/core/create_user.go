package core

import (
	"context"
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

	if err = req.Validate(); err != nil {
		return nil, err
	}

	var createdId int
	err = c.db.QueryRow(`insert into users(role, name, surname, patronymic, tel, email, commission, password)
							 values($1, $2, $3, $4, $5, $6, $7, $8)
							 returning id`,
		req.Role,
		req.Name,
		req.Surname,
		req.Patronymic,
		req.Telephone,
		req.Email,
		req.Commission,
		req.Password,
	).Scan(&createdId)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetUser(ctx, createdId)
}
