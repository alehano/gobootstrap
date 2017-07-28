package example

import (
	"github.com/alehano/gobootstrap/sys/pubsub"
	"github.com/alehano/gobootstrap/config"
)

func init() {

	pubsub.Subscribe(config.ExampleCreatedMsg, func(data interface{}) {
		//id := data.(int)
		// ...
	})

}
