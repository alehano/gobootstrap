package models

import (
	"errors"
	"database/sql"
)

// DB agnostic errors
var ErrNotExists = errors.New("Not exists")

// Get rid of Postgres error type
func WrapPostgresErr(err error) error {
	if err == sql.ErrNoRows {
		err = ErrNotExists
	}
	return err
}