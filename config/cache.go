package config

// Centralized storage for cache keys
var CacheKeys = cacheKeys{}

type cacheKeys struct{}

func (c cacheKeys) Example(id int) string {
	return "TODO"
}