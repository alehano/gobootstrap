/*
Wrapper for Logrus

log.WithFields(log.Fields{
	"key": "val",
	"key2": "val2",
}).Info("Index page")

*/
package log

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})

	// Add your hook ...
	// https://github.com/sirupsen/logrus
}
