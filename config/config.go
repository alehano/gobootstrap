/*
Config package contains different typed parameters accessible by .Get(),
as well as some structures like cache keys generator.
Config file can be load by either set environment variable "APP_CONFIG" with
full path to .yml config file or by putting config.yml to the app working directory.

Usage:
    `config.Get().SomeParam`

 */
package config

var Version = "0.0.1"

type cfg struct {
	Debug             bool
	Port              int
	ProjectPath       string `yaml:"project_path"`
	AdminLogin        string `yaml:"admin_login"`
	AdminPasswordHash string `yaml:"admin_password_hash"`
	MemcachedAddr     []string `yaml:"memcached_addr"`
	PostgresHost      string `yaml:"postgres_host"`
	PostgresDatabase  string `yaml:"postgres_database"`
	PostgresUser      string `yaml:"postgres_user"`
	PostgresPassword  string `yaml:"postgres_password"`
	PostgresSSLMode   string `yaml:"postgres_ssl_mode"` // e.g. "disable"

	// ... add more
}

func Defaults() cfg {
	return cfg{
		Debug: false,
		Port:  8000,

		// ... add more
	}
}
