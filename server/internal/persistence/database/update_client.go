package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) UpdateClient(ctx context.Context, req *domain.UpdateClientRequest) error {
	var (
		args   = make([]any, 0)
		values []string
	)

	if req.Name != nil {
		args = append(args, *req.Name)
		values = append(values, "name = ?")
	}

	if req.Surname != nil {
		args = append(args, *req.Surname)
		values = append(values, "surname = ?")
	}

	if req.Patronymic != nil {
		args = append(args, *req.Patronymic)
		values = append(values, "patronymic = ?")
	}

	if req.Email != nil {
		args = append(args, *req.Email)
		values = append(values, "email = ?")
	}

	if req.Telephone != nil {
		args = append(args, *req.Telephone)
		values = append(values, "tel = ?")
	}

	q := strings.Join(values, ",")
	query := fmt.Sprintf("update clients set %s where id = ?", q)
	args = append(args, req.ClientId)

	_, err := d.Conn.ExecContext(ctx, query, args...)
	return err
}
