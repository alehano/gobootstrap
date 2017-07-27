package memcached

import (
	"testing"
	"reflect"
)

func TestGetSetObj(t *testing.T) {
	res := []string{}
	firstCall := []string{"first", "call"}
	secondCall := []string{"second", "call"}

	err := GetSetObj("test", &res, func() (interface{}, error) {
		return firstCall, nil
	}, Expiration30d, SnappyCompressFlag)
	if err != nil {
		t.Error(err)
	}

	err = GetSetObj("test", &res, func() (interface{}, error) {
		return secondCall, nil
	}, Expiration30d, SnappyCompressFlag)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(res, firstCall) {
		t.Errorf("Not equal %v != %v", res, firstCall)
	}
}

func TestGetSetObjInts(t *testing.T) {
	res := []int{}

	err := GetSetObj("test", &res, func() (interface{}, error) {
		return []int{1, 2}, nil
	}, Expiration30d)
	if err != nil {
		t.Error(err)
	}

	t.Logf("RES: %v", res)

	if len(res) != 2 {
		t.Errorf("Len wrong %v", res)
	}

	if !reflect.DeepEqual(res, []int{1, 2}) {
		t.Errorf("Not equal %v != %v", res, []int{1, 2})
	}
}
