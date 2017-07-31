package config

import (
	"log"
	"github.com/alehano/gobootstrap/sys/cmd"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

func init() {

	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show version",
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Version: %s\n", Version)
		},
	})

	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "admin_pwd [password]",
		Short: "Generate Admin password hash",
		Long: "Generate Admin password hash (BCrypt) for using in config file.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				hash, err := bcrypt.GenerateFromPassword([]byte(args[0]), bcrypt.DefaultCost)
				if err != nil {
					return err
				} else {
					log.Printf("Password Hash: %s\n", hash)
				}
			} else {
				log.Println("Error: password string wasn't provided")
			}
			return nil
		},
	})


	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "init_config",
		Short: "Init default config file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return CreateDefaultConfigFile()
		},
	})

	//cli.RegisterCLI(ConfigInitCLI,
	//	fmt.Sprintf("Create default config file (%s)", defaultFilename), CreateDefaultConfigFile)
	//
	//cli.RegisterCLIWithArgs(GenAdminPwd, "[string] Generate Admin password", func(args ...string) error {
	//	fmt.Sprintf("Args: %v", args)
	//	return nil
	//})

}
