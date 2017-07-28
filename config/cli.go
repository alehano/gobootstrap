package config

import (
	"github.com/alehano/gobootstrap/sys/cli"
	"fmt"
)

func init() {
	cli.RegisterCLI(VersionCLI, "Get App version", func() error {
		fmt.Println(Version)
		return nil
	})

	cli.RegisterCLI(ConfigInitCLI,
		fmt.Sprintf("Create default config file (%s)", defaultFilename), CreateDefaultConfigFile)
}
