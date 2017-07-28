package pubsub

import (
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	t1pass := 0
	t2pass := 0

	testFn1 := func(msg interface{}) {
		if msg.(string) == "T1" {
			t1pass++
		}
	}

	testFn11 := func(msg interface{}) {
		if msg.(string) == "T1" {
			t1pass++
		}
	}

	testFn2 := func(msg interface{}) {
		if msg.(string) == "T2" {
			t2pass++
		}
	}

	Subscribe("test1", testFn1)
	Subscribe("test1", testFn11)
	Subscribe("test2", testFn2)

	Publish("test1", "T1")
	Publish("test2", "T2")

	time.Sleep(100*time.Millisecond)

	if t1pass != 2 || t2pass != 1 {
		t.Errorf("Not Pass T1:%v, T2:%v", t1pass, t2pass)
	}
}
