/*
Centralized storage for cache keys
 */
package config

var CacheKeys = cacheKeys{}

type cacheKeys struct{}

func (c cacheKeys) Foo(bar string) string {
	return "TODO"
}

func (c cacheKeys) Foo2() string { return "TODO" }
