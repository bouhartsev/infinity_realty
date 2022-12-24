package core

import (
	"github.com/jackc/pgx"
	"go.uber.org/zap"

	"github.com/bouhartsev/infinity_realty/server/internal/config"
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
