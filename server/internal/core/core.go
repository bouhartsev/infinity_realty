package core

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/bouhartsev/infinity_realty/internal/config"
	"github.com/bouhartsev/infinity_realty/internal/database"
	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

type Core struct {
	logger *zap.Logger
	db     *database.Database
	cfg    *config.Config
}

func NewCore(l *zap.Logger, db *pgxpool.Pool, c *config.Config) *Core {
	return &Core{
		logger: l,
		db:     &database.Database{Conn: db},
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
