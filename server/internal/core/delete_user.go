package core

import (
	"context"
	"errors"

	"github.com/bouhartsev/infinity_realty/internal/database"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (c *Core) DeleteUser(ctx context.Context, id int) error {
	user, err := c.getUserFromCtx(ctx)
	if err != nil {
		return err
	}

	if user.Role != 1 {
		return errdomain.ForbiddenError
	}

	_, err = c.GetUser(ctx, id)
	if err != nil {
		return err
	}

	err = c.db.CheckUserDependencies(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrUserHasNeeds) {
			return errdomain.UserHasNeedsError
		}
		return errdomain.NewInternalError(err.Error())
	}

	err = c.db.DeleteUser(ctx, id)
	if err != nil {
		return errdomain.NewInternalError(err.Error())
	}

	return nil
}
