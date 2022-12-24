package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) UpdateUser(ctx context.Context, req *domain.UpdateUserRequest) error {
	var (
		argId  = 1
		args   = make([]any, 0)
		values []string
	)

	if req.Name != nil {
		args = append(args, *req.Name)
		values = append(values, fmt.Sprintf("name = $%d", argId))
		argId++
	}

	if req.Surname != nil {
		args = append(args, *req.Surname)
		values = append(values, fmt.Sprintf("surname = $%d", argId))
		argId++
	}

	if req.Patronymic != nil {
		args = append(args, *req.Patronymic)
		values = append(values, fmt.Sprintf("patronymic = $%d", argId))
		argId++
	}

	if req.Email != nil {
		args = append(args, *req.Email)
		values = append(values, fmt.Sprintf("email = $%d", argId))
		argId++
	}

	if req.Telephone != nil {
		args = append(args, *req.Telephone)
		values = append(values, fmt.Sprintf("tel = $%d", argId))
		argId++
	}

	if req.Commission != nil {
		args = append(args, *req.Commission)
		values = append(values, fmt.Sprintf("commission = $%d", argId))
		argId++
	}

	if req.Password != nil {
		args = append(args, *req.Password)
		values = append(values, fmt.Sprintf("password = $%d", argId))
		argId++
	}

	q := strings.Join(values, ",")
	query := fmt.Sprintf("update users set %s where id = $%d", q, argId)
	args = append(args, req.UserId)

	_, err := d.Conn.Exec(ctx, query, args...)
	return err
}
