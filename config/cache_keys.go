package config

import "fmt"

// Centralized storage for cache keys
var CacheKeys = cacheKeys{}

type cacheKeys struct{}

func (c cacheKeys) ExampleGet(id int) string {
	return fmt.Sprintf("example.get.%d", id)
}

// .. add more