package db

import (
	"time"
)

const (
	connTimeToReconnect = 5 * time.Second
	connMaxTries        = 10
)

// Waiting some time (increasing) to reconnect,
// After n tries fails
func ReconnectCounter(tries ...int) int {
	try := 1
	if len(tries) > 0 {
		try = tries[0] + 1
		if try > connMaxTries {
			panic("Max DB connections failed")
		}
	}
	time.Sleep(time.Duration(try) * connTimeToReconnect)
	return try
}
