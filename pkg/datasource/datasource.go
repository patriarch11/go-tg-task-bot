package datasource

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Datasource contains Postgres client for repositories
type Datasource struct {
	*pgxpool.Pool
}

// New function creates new Postgres datasource
func New(ctx context.Context, config *Config) (*Datasource, error) {
	var err error
	s := new(Datasource)
	connConfig, err := pgxpool.ParseConfig(config.URL)
	if err != nil {
		return nil, err
	}
	connConfig.ConnConfig.PreferSimpleProtocol = config.PreferSimpleProtocol
	connConfig.MaxConns = config.MaxConnections
	connConfig.MinConns = config.IdleConnections
	s.Pool, err = pgxpool.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, err
	}
	return s, nil
}
