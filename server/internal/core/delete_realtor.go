package core

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) DeleteRealtor(ctx context.Context, id int) error {
	_, err := c.GetRealtor(ctx, id)
	if err != nil {
		return err
	}

	err = c.db.DeleteRealtor(ctx, id)
	if err != nil {
		return errdomain.NewInternalError(err.Error())
	}

	return nil
}
