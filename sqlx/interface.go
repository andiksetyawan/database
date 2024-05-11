//go:generate go run github.com/vektra/mockery/v2@v2.43.0 --name "DB|Q"
package sqlx

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DB interface {
	GetMaster() *sqlx.DB
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
}

// Q querier
type Q interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}
