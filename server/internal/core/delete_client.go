package core

import (
	"context"
	"errors"

	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
	"github.com/bouhartsev/infinity_realty/internal/persistence/database"
)

func (c *Core) DeleteClient(ctx context.Context, id int) error {
	_, err := c.GetClient(ctx, id)
	if err != nil {
		return err
	}

	err = c.db.CheckClientDependencies(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrClientHasNeeds) {
			return errdomain.ClientHasNeedsError
		}
		return errdomain.NewInternalError(err.Error())
	}

	err = c.db.DeleteClient(ctx, id)
	if err != nil {
		return errdomain.NewInternalError(err.Error())
	}

	return nil
}
