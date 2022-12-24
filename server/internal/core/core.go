package core

import (
	"context"
	"github.com/bouhartsev/infinity_realty/internal/config"
	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
	"github.com/jackc/pgx"
	"go.uber.org/zap"
)

type Core struct {
	logger *zap.Logger
	db     *pgx.ConnPool
	cfg    *config.Config
}

func NewCore(l *zap.Logger, db *pgx.ConnPool, c *config.Config) *Core {
	return &Core{
		logger: l,
		db:     db,
		cfg:    c,
	}
}

func (c *Core) getUserFromCtx(ctx context.Context) (*domain.User, error) {
	val := ctx.Value("user")
	user, ok := val.(domain.User)
	if !ok {
		return nil, errdomain.NewInternalError("token has not been parsed to context")
	}
	return &user, nil
}
