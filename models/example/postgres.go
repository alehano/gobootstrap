package example

import (
	"github.com/alehano/gobootstrap/sys/db/postgres"
	"github.com/alehano/gobootstrap/models/common"
	"github.com/jmoiron/sqlx"
	"errors"
)

// Implementation of Storage interface
func NewPostgresStorage() ExampleStorage {
	s := postgresStorage{
		db:    postgres.GetDB(),
		table: tableName,
	}
	return s
}

const (
	tableName = "example"
)

type postgresStorage struct {
	db    *sqlx.DB
	table string
}

func (s postgresStorage) DBInit() error {
	s.db.MustExec(`
	CREATE TABLE $1 (
	    id SERIAL PRIMARY KEY,
	    title VARCHAR(100) NOT NULL,
	    updated_at TIMESTAMP,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`, s.table)
	return common.WrapPostgresErr(errors.New("TODO"))
}

func (s postgresStorage) Create(item ExampleModel) (int, error) {
	res, err := s.db.Exec("INSERT INTO $1 (title, updated_at) VALUES ($2, NOW())",
		s.table, item.Title)
	if err != nil {
		return 0, common.WrapPostgresErr(err)
	}
	newId, err := res.LastInsertId()
	return int(newId), common.WrapPostgresErr(err)
}

func (s postgresStorage) Get(id int) (ExampleModel, error) {
	res := ExampleModel{}
	err := s.db.Get(&res, "SELECT * FROM $1 WHERE id=$2", s.table, id)
	return res, common.WrapPostgresErr(err)
}

func (s postgresStorage) Update(item ExampleModel) error {
	return common.WrapPostgresErr(errors.New("TODO"))
}

func (s postgresStorage) Delete(id int) error {
	return common.WrapPostgresErr(errors.New("TODO"))
}
