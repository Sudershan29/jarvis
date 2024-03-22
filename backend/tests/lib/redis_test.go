
package redis_test

import (
	"errors"
	"testing"
	"backend/src/lib"
)


func TestRedisStore(t *testing.T) {
	store  := lib.ConnectToRedis()
	err := store.Store("hello", "1")
    if err != nil {
        t.Errorf("Cannot store in redis : %s", err.Error())
    }

	val, err := store.Get("hello")
	if err != nil {
        t.Errorf("Cannot get from redis : %s", err.Error())
	}

	if val != "1" {
		t.Errorf("Value did not match : %s", errors.New("Redis('hello') != 1"))
	}
}