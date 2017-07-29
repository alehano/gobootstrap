/*
Config contains different parameters as well as some lists of names used in other App parts,
like cache keys, CLI commands or PubSub messages.
Config file can be load by set environment variable "APP_CONFIG" with
full path to .yml config file. Or by putting config.yml to the app root directory.

Usage:
    `config.Get().SomeParam`

 */
package config

var Version = "0.0.1"

type cfg struct {
	Debug       bool
	Port        int
	ProjectPath string `yaml:"project_path"`

	MemcachedAddr    []string `yaml:"memcached_addr"`
	PostgresHost     string `yaml:"postgres_host"`
	PostgresDatabase string `yaml:"postgres_database"`
	PostgresUser     string `yaml:"postgres_user"`
	PostgresPassword string `yaml:"postgres_password"`
	PostgresSSLMode  string `yaml:"postgres_ssl_mode"` // e.g. "disable"

	// ... add more
}

func GetDefault() cfg {
	return cfg{
		Debug: false,
		Port:  8000,

		// ... add more
	}
}
