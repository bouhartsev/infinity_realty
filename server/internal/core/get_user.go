package core

import (
	"context"
	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) GetUser(ctx context.Context, id int) (*domain.GetUserResponse, error) {
	var user domain.User

	err := c.db.QueryRow(`select id, role, name, surname, patronymic, tel, email, commission
                              from users
                              where id = $1`,
		id).Scan(
		&user.Id,
		&user.Role,
		&user.Name,
		&user.Surname,
		&user.Patronymic,
		&user.Telephone,
		&user.Email,
		&user.Commission,
	)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return &domain.GetUserResponse{User: &user}, nil
}