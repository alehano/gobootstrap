package example

import "github.com/alehano/gobootstrap/sys/db"

// Instance of Manager
var Man Manager = NewExampleMan(NewPostgresStorage())

func init() {
	db.RegisterInitter(Man.storage)
}

