/*
Wrapper for Memcached
  */
package memcached

import (
	"github.com/alehano/gobootstrap/config"
	"github.com/alehano/gohelpers/compress"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

const (
	NothingFlag        uint32 = 0
	SnappyCompressFlag uint32 = 1
	GzipCompressFlag   uint32 = 2
)

var (
	// Inited memcached Client
	mc = memcache.New(config.Get().MemcachedAddr...)

	compressMethods = map[uint32]string{
		SnappyCompressFlag: "snappy",
		GzipCompressFlag:   "gz",
	}

	// Predefined expiration time
	NoExpiration  = 0
	Expiration15m = int32((15 * time.Minute).Seconds())
	Expiration1h  = int32((time.Hour).Seconds())
	Expiration24h = int32((24 * time.Hour).Seconds())
	Expiration3d  = int32((3 * 24 * time.Hour).Seconds())
	Expiration30d = int32((30 * 24 * time.Hour).Seconds())
)

// Using when need expiration > 30d
func SetExpiration(dur time.Duration) int32 {
	return int32(time.Now().Add(dur).Unix())
}

// GetSet functions. Helpers to get value from cache or calculate and set it.

// GetSetObj gets object from cache and save it to dest pointer.
// If object not exists, get it from getFn and save to a cache.
func GetSetObj(key string, dest interface{}, getFn func() (interface{}, error), expire int32, flag ...uint32) error {
	found, err := GetObj(key, dest)
	if err != nil {
		return err
	}
	if !found {
		newObj, err := getFn()
		if err != nil {
			return err
		}
		err = SetObj(key, newObj, expire, flag...)
		if err != nil {
			return err
		}
		found, err := GetObj(key, dest)
		if err != nil {
			return err
		}
		if !found {
			return errors.New(fmt.Sprintf("GetSetObj not found after set. Key: %s", key))
		}
	}
	return nil
}

func GetSetString(key string, getFn func() (string, error), expire int32, flag ...uint32) (string, error) {
	res, found, err := GetString(key)
	if err != nil {
		return res, err
	}
	if !found {
		newRes, err := getFn()
		if err != nil {
			return res, err
		}
		err = SetString(key, newRes, expire, flag...)
		return newRes, err
	}
	return res, nil
}

func GetSetInt(key string, getFn func() (int, error), expire int32, flag ...uint32) (int, error) {
	res, found, err := GetInt(key)
	if err != nil {
		return res, err
	}
	if !found {
		newRes, err := getFn()
		if err != nil {
			return res, err
		}
		err = SetInt(key, newRes, expire, flag...)
		return newRes, err
	}
	return res, nil
}

func GetSetBool(key string, getFn func() (bool, error), expire int32, flag ...uint32) (bool, error) {
	res := false
	num, found, err := GetInt(key)
	if err != nil {
		return res, err
	}
	if !found {
		newRes, err := getFn()
		if err != nil {
			return newRes, err
		}
		if newRes {
			num = 1
		} else {
			num = 0
		}
		err = SetInt(key, num, expire, flag...)
		return newRes, err
	}
	if num == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

// Set functions

func SetBytes(key string, value []byte, expire int32, flag ...uint32) error {
	if len(value) == 0 {
		return nil
	}
	var err error
	if len(flag) > 0 {
		// Compress
		if method, ok := compressMethods[flag[0]]; ok {
			value, err = compress.Encode(value, method)
			if err != nil {
				return errors.New(fmt.Sprintf("Cache compress. Method: %s, Key: %s, Value: %+v, Error: %s", method, key, value, err))
			}
		}
		err = mc.Set(&memcache.Item{Key: key, Value: value, Expiration: expire, Flags: flag[0]})
	} else {
		err = mc.Set(&memcache.Item{Key: key, Value: value, Expiration: expire, Flags: NothingFlag})
	}
	if err != nil {
		return errors.New(fmt.Sprintf("Cache Set Key: %s, Err: %s", key, err))
	}
	return nil
}

func SetString(key, value string, expire int32, flag ...uint32) error {
	return SetBytes(key, []byte(value), expire, flag...)
}

func SetInt(key string, value int, expire int32, flag ...uint32) error {
	return SetBytes(key, []byte(strconv.Itoa(value)), expire, flag...)
}

func SetFloat(key string, value float64, expire int32, flag ...uint32) error {
	return SetBytes(key, []byte(strconv.FormatFloat(value, 'E', -1, 64)), expire, flag...)
}

func SetObj(key string, obj interface{}, expire int32, flag ...uint32) error {
	var err error
	objBytes, err := json.Marshal(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("Cache Marshal Key: %s, Obj: %+v, Err: %s", key, obj, err))
	}
	return SetBytes(key, objBytes, expire, flag...)
}

// Get functions

func GetBytes(key string) ([]byte, bool, error) {
	item, err := mc.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return []byte{}, false, nil
		}
		return []byte{}, false, errors.New(fmt.Sprintf("Cache Get Key: %s, Err: %s", key, err))
	}
	if method, ok := compressMethods[item.Flags]; ok {
		item.Value, err = compress.Decode(item.Value, method)
		if err != nil {
			return []byte{}, false, errors.New(fmt.Sprintf("Cache Decode. Method: %s, Key: %s, Value: %+v, Err: %s",
				method, key, item.Value, err))
		}
	}
	return item.Value, true, nil
}

// Returns: value, exists, error
func GetString(key string) (string, bool, error) {
	data, exists, err := GetBytes(key)
	if err != nil {
		return "", false, err
	}
	if !exists {
		return "", false, nil
	}
	return string(data), true, nil
}

func GetInt(key string) (int, bool, error) {
	data, exists, err := GetBytes(key)
	if err != nil {
		return 0, false, err
	}
	if !exists {
		return 0, false, nil
	}
	res, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, false, errors.New(fmt.Sprintf("Cache Atoi Key: %s, Value: %+v, Err: %s",
			key, data, err))
	}
	return res, true, nil
}

func GetFloat(key string) (float64, bool, error) {
	data, exists, err := GetBytes(key)
	if err != nil {
		return 0, false, err
	}
	if !exists {
		return 0, false, nil
	}
	res, err := strconv.ParseFloat(string(data), 10)
	if err != nil {
		return 0, false, errors.New(fmt.Sprintf("Cache ParseFloat Key: %s, Value: %+v, Err: %s",
			key, data, err))
	}
	return res, true, nil
}

// GetObj decodes object to dest pointer
// Returns false if not exists or error
func GetObj(key string, dest interface{}) (bool, error) {
	data, exists, err := GetBytes(key)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}
	err = json.Unmarshal(data, dest)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Cache Unmarshal Key: %s, Value: %+v, Err: %s",
			key, data, err))
	}
	return true, nil
}

// Delete removes item form a cache
func Delete(key string) error {
	err := mc.Delete(key)
	if err != nil && err == memcache.ErrCacheMiss {
		return nil
	}
	return err
}
