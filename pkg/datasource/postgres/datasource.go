package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// New function creates new Postgres datasource
func New(ctx context.Context, config *Config) (*pgxpool.Pool, error) {
	connConfig, err := pgxpool.ParseConfig(config.URL)
	if err != nil {
		return nil, err
	}
	connConfig.ConnConfig.Logger = NewLogger(logrus.StandardLogger())
	connConfig.ConnConfig.PreferSimpleProtocol = config.PreferSimpleProtocol
	connConfig.MaxConns = config.MaxConnections
	connConfig.MinConns = config.IdleConnections
	return pgxpool.ConnectConfig(ctx, connConfig)
}
