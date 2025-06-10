package zcache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGet_Success(t *testing.T) {
	cache := NewCache[string, int]()

	err := cache.Set("test key 1", 100)
	assert.Nil(t, err, "Set -- Should add the value and key to the cache")

	value, found := cache.Get("test key 1")
	assert.True(t, found, "Get -- Should find the key in the cache")
	assert.Equal(t, 100, value, "Get -- Should get the correct value associated with the key")
}

func TestGet_EntryNotFound(t *testing.T) {
	cache := NewCache[string, int]()

	err := cache.Set("test key 1", 100)
	assert.Nil(t, err, "Set -- Should add the value and key to the cache")

	_, found := cache.Get("test key 2")
	assert.False(t, found, "Get -- Should not find the given key in the cache")
}

func TestGet_CacheEmpty(t *testing.T) {
	cache := NewCache[string, int]()

	_, found := cache.Get("test key 1")
	assert.False(t, found, "Get -- Cache should be empty and Get should return false")
}

func TestSet_KeyAlreadyExists(t *testing.T) {
	cache := NewCache[string, int]()

	err := cache.Set("test key 1", 100)
	assert.Nil(t, err, "Set -- Should add the value and key to the cache")

	err = cache.Set("test key 1", 200)
	assert.NotNil(
		t,
		err,
		"Set -- Should return an error when trying to set a key that already exists",
	)
}
