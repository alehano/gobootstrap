package example

import (
	"database/sql"
	"github.com/alehano/gobootstrap/sys/db/postgres"
	"github.com/alehano/gobootstrap/models/common"
	"errors"
)

func NewPostgresStorage() ExampleStorage {
	s := postgresStorage{db: postgres.GetDB()}
	return s
}

// Implementation of Storage interface
type postgresStorage struct {
	db *sql.DB
}

func (s postgresStorage) DBInit() error {
	return common.WrapPostgresErr(errors.New("TODO"))
}

func (s postgresStorage) Create(item ExampleModel) (int, error) {
	return 0, common.WrapPostgresErr(errors.New("TODO"))
}

func (s postgresStorage) Get(id int) (ExampleModel, error) {
	return ExampleModel{}, common.WrapPostgresErr(errors.New("TODO"))
}

func (s postgresStorage) Update(item ExampleModel) error {
	return common.WrapPostgresErr(errors.New("TODO"))
}

func (s postgresStorage) Delete(id int) error {
	return common.WrapPostgresErr(errors.New("TODO"))
}
