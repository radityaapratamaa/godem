package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Master interface {
	Exec(query string, args ...interface{}) (sql.Result, error)

	// ExecContext use master database to exec query
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// Begin transaction on master DB
	Begin() (*sql.Tx, error)

	// BeginTx begins transaction on master DB
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	// Rebind a query from the default bindtype (QUESTION) to the target bindtype.
	Rebind(sql string) string

	// NamedExec do named exec on master DB
	NamedExec(query string, arg interface{}) (sql.Result, error)

	// NamedExecContext do named exec on master DB
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)

	// BindNamed do BindNamed on master DB
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
}

// Follower defines operation that will be executed to follower DB
type Follower interface {
	// Get from follower database
	Get(dest interface{}, query string, args ...interface{}) error

	// Select from follower database
	Select(dest interface{}, query string, args ...interface{}) error

	// Query from follower database
	Query(query string, args ...interface{}) (*sql.Rows, error)

	// QueryRow executes QueryRow against follower DB
	QueryRow(query string, args ...interface{}) *sql.Row

	// NamedQuery do named query on follower DB
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)

	// GetContext from sql database
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// SelectContext from sql database
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// QueryContext from sql database
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	// QueryRowContext from sql database
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row

	// QueryxContext queries the database and returns an *sqlx.Rows. Any placeholder parameters are replaced with supplied args.
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)

	// QueryRowxContext queries the database and returns an *sqlx.Row. Any placeholder parameters are replaced with supplied args.
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row

	// NamedQueryContext do named query on follower DB
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
}
