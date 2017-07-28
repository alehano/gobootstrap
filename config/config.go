/*
Config contains different parameters as well as some lists of names used in other App parts,
like cache keys, CLI commands or pubsub messages.
Config file can be load by set environment variable "APP_CONFIG" with
full path to yaml config file. Or by placing config.yml to the app root directory.

Usage:
    `config.Get().SomeParam`

 */
package config

var Version = "0.0.1"

type cfg struct {
	Debug       bool
	Port        int
	ProjectPath string

	MemcachedAddr    []string
	PostgresHost     string
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	PostgresSSLMode  string // e.g. "disable"

	// ... add more
}

func GetDefault() cfg {
	return cfg{
		Debug: false,
		Port:  8000,

		// ... add more
	}
}
