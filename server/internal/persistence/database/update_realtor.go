package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) UpdateRealtor(ctx context.Context, req *domain.UpdateRealtorRequest) error {
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

	if req.Commission != nil {
		args = append(args, *req.Commission)
		values = append(values, "commission = ?")
	}

	q := strings.Join(values, ",")
	query := fmt.Sprintf("update realtors set %s where id = ?", q)
	args = append(args, req.RealtorId)

	_, err := d.Conn.ExecContext(ctx, query, args...)
	return err
}
