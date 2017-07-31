package models

import (
	"github.com/alehano/gobootstrap/sys/cmd"
	"github.com/alehano/gobootstrap/sys/db"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "init_db",
		Short: "Init all DB",
		Long: "Init all DB tables with DBInitter interface being registered in sys/db",
		RunE: func(cmd *cobra.Command, args []string) error {
			return db.InitAllDBs()
		},
	})
}

