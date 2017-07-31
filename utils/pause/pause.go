/*
Pause sleeps on each step for some time (increasing in linear progression).
 */
package pause

import (
	"time"
)

// New creates Pause object.
// max - maximum tries, interval - time to sleep (n * interval).
func New(max int, interval time.Duration) *p {
	return &p{max: max, interval: interval}
}

type p struct {
	max      int
	interval time.Duration
	n        int
}

// Do returns false if max tries is reached.
// Instead returns true and sleeps some time (n * interval).
func (r *p) Do() bool {
	if r.n >= r.max {
		return false
	}
	time.Sleep(time.Duration(r.n) * r.interval)
	r.n++
	return true
}
