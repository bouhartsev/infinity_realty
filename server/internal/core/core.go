package core

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/bouhartsev/infinity_realty/internal/config"
	"github.com/bouhartsev/infinity_realty/internal/database"
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
