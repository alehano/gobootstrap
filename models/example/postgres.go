package example

import (
	"github.com/alehano/gobootstrap/sys/db/postgres"
	"github.com/alehano/gobootstrap/models"
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
	return models.WrapPostgresErr(errors.New("TODO"))
}

func (s postgresStorage) Create(item ExampleModel) (int, error) {
	res, err := s.db.Exec("INSERT INTO $1 (title, updated_at) VALUES ($2, NOW())",
		s.table, item.Title)
	if err != nil {
		return 0, models.WrapPostgresErr(err)
	}
	newId, err := res.LastInsertId()
	return int(newId), models.WrapPostgresErr(err)
}

func (s postgresStorage) Get(id int) (ExampleModel, error) {
	res := ExampleModel{}
	err := s.db.Get(&res, "SELECT * FROM $1 WHERE id=$2", s.table, id)
	return res, models.WrapPostgresErr(err)
}

func (s postgresStorage) Update(item ExampleModel) error {
	return models.WrapPostgresErr(errors.New("TODO"))
}

func (s postgresStorage) Delete(id int) error {
	return models.WrapPostgresErr(errors.New("TODO"))
}
