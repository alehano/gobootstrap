/*
Publish-Subscriber local service
Can be replaced by distributed one (e.g. http://nsq.io) in case of multi-node configuration
 */
package pubsub

import "time"

var pubSubStore = map[string][]func(interface{}){}

// Subscribe - register func to execute when message with this topic occur.
// Func executes in async way.
func Subscribe(topic string, fn func(interface{})) {
	pubSubStore[topic] = append(pubSubStore[topic], fn)
}

func SubscribeMultiple(topics []string, fn func(interface{})) {
	for _, topic := range topics {
		Subscribe(topic, fn)
	}
}

// Publish - async send messages with given topic
// Runs each fn in goroutine
func Publish(topic string, msg interface{}, delay ...time.Duration) {
	fns, ok := pubSubStore[topic]
	if ok {
		for _, fn := range fns {
			go func(f func(interface{})) {
				if len(delay) > 0 {
					time.Sleep(delay[0])
				}
				f(msg)
			}(fn)
		}
	}
}

// Publish in a sync manner
func PublishSync(topic string, msg interface{}) {
	fns, ok := pubSubStore[topic]
	if ok {
		for _, fn := range fns {
			fn(msg)
		}
	}
}
