package config

import "fmt"

// Centralized storage for cache keys
var CacheKeys = cacheKeys{}

type cacheKeys struct{}

func (c cacheKeys) ExampleGet(id int) string {
	return fmt.Sprintf("example.get.%d", id)
}

func (c cacheKeys) AuthThrottle(login string) string {
	// TODO: hash murmur login
	return fmt.Sprintf("auth.thr.%s", login)
}

// .. add more