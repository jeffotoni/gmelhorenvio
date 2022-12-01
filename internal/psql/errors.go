package psql

import "errors"

var (
	ErrDatabaseDown = errors.New("database down")
)
