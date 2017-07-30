package models

import (
	"github.com/alehano/gobootstrap/sys/cmd"
	"github.com/alehano/gobootstrap/sys/db"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "init_db",
		Short: "Init all DB",
		Long: "Init all DB tables with DBInitter interface being registered in sys/db",
		Run: func(cmd *cobra.Command, args []string) {
			err := db.InitAllDBs()
			if err != nil {
				log.Println(err)
			}
		},
	})
}

