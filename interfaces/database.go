package interfaces

import (
	"database/sql"
	"julo/domain"
)

// IDatabase is interface for database
type Database interface {
	ConnectDB(dbAcc *domain.DBAccount)
	Close()

	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	// DriverName() string

	// Begin() (IDBTx, error)
	In(query string, params ...interface{}) (string, []interface{}, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	// QueryRow(query string, args ...interface{}) *sql.Row
	// QueryRowSqlx(query string, args ...interface{}) *sqlx.Row
	// QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
