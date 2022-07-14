package storage

import (
	"database/sql"
	"fmt"
	"product-service/internal/config"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewPostgreStorage(cfg *config.Config) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.Database)
	conn, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
