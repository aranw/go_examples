package pbt_test

import (
	"testing"
	"testing/quick"

	. "github.com/aranw/go_examples/pbt"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	set := func(key string, value string) interface{} {
		cache.Add(key, value)

		return value
	}
	get := func(key string, value string) interface{} {
		v, _ := cache.Get(key)
		return v
	}
	if err := quick.CheckEqual(set, get, nil); err != nil {
		t.Error(err)
	}
}

func TestCacheErrorsOnUnknownKey(t *testing.T) {
	cache := NewCache()

	_, err := cache.Get("blah")

	if err == nil {
		t.Error("Expected an error")
	}
}
