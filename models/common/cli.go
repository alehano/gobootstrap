package common

import (
	"github.com/alehano/gobootstrap/sys/cli"
	"github.com/alehano/gobootstrap/sys/db"
	"github.com/alehano/gobootstrap/config"
)

func init() {
	cli.RegisterCLI(config.DBInitAllCLI, "Init all DB data", db.InitAllDBs)
}

