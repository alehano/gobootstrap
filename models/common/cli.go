package common

import (
	"github.com/alehano/gobootstrap/sys/cli"
	"github.com/alehano/gobootstrap/sys/db"
)

func init() {
	cli.RegisterCLI("db_init_all", "Init all DB data", db.InitAllDBs)
}

