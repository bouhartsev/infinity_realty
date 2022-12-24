package core

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) SignIn(ctx context.Context, request *domain.SignInRequest) (*domain.SignInResponse, error) {
	var user domain.User

	q := "select id, role, name, surname, patronymic, tel, email, commission from users where (email = $1 or telephone = $1) and password = $2"

	row := c.db.QueryRow(ctx, q, request.Login, request.Password)

	err := row.Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.InvalidCredentialsError
		}
		return nil, errdomain.NewInternalError(err.Error())
	}

	claims := domain.AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: user,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(c.cfg.TokenKey))

	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return &domain.SignInResponse{Token: token, User: user}, nil
}
