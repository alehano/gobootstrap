package example

import (
	"github.com/alehano/gobootstrap/sys/pubsub"
	"github.com/alehano/gobootstrap/config"
	"github.com/alehano/gobootstrap/sys/memcache"
)

func init() {

	pubsub.Subscribe(config.ExampleCreatedMsg, func(data interface{}) {
		//id := data.(int)
		// ...
	})

	// Invalidate cache
	pubsub.SubscribeMultiple([]string{config.ExampleUpdateMsg, config.ExampleDeleteMsg},
		func(data interface{}) {
			memcached.Delete(config.CacheKeys.ExampleGet(data.(int)))
		})

}
