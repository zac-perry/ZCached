package server

import (
	"testing"
	"time"
)

/*
--------------------------

# Get Function Tests

--------------------------
*/
func TestGet_ValueExists(t *testing.T)       {}
func TestGet_ValueDoesNotExist(t *testing.T) {}
func TestGet_ValueIsExpired(t *testing.T)    {}

/*
--------------------------

# Put Function Tests

--------------------------
*/
func TestPut_Success(t *testing.T)          {}
func TestPut_KeyAlreadyExists(t *testing.T) {}
func TestPut_CacheIsFull(t *testing.T)      {}

/*
--------------------------

# IsExpired Function Tests

--------------------------
*/
func TestIsExpired_WhenExpired(t *testing.T) {
	expire := time.Second
	cacheEntry := Entry{
		key:        "TestKey",
		value:      100,
		createdAt:  time.Now().Add(-2 * time.Second),
		expireTime: &expire,
	}

	if !cacheEntry.isExpired() {
		t.Error("Cache entry should be expired")
	}
}

func TestIsExpired_WhenNotExpired(t *testing.T) {
	expire := time.Hour
	cacheEntry := Entry{
		key:        "TestKey",
		value:      100,
		createdAt:  time.Now(),
		expireTime: &expire,
	}

	if cacheEntry.isExpired() {
		t.Error("Cache entry should NOT be expired")
	}
}

func TestIsExpired_WhenNeverExpired(t *testing.T) {
	cacheEntry := Entry{
		key:        "TestKey",
		value:      100,
		createdAt:  time.Now(),
		expireTime: nil,
	}

	if cacheEntry.isExpired() {
		t.Error("Cache entry should NOT be expired and should NEVER expire")
	}
}

// Push front

/*
--------------------------

# RemoveEntry Function Tests

--------------------------
*/
func TestRemoveEntry_RemoveLastEntry(t *testing.T) {
	cacheEntry1 := Entry{
		key:   "1",
		value: 100,
	}

	cacheEntry2 := Entry{
		key:   "2",
		value: 200,
	}

	cache := NewCache(2)
	cache.head = &cacheEntry1
	cache.tail = &cacheEntry2
	cache.head.Next = cache.tail
	cache.tail.Prev = cache.head
	cache.entries[cache.head.key] = cache.head
	cache.entries[cache.tail.key] = cache.tail

	cache.RemoveLRUEntry()

	if len(cache.entries) != 1 {
		t.Error("Error removing the last entry of the cache")
	}

	// in this case, head and tail should be the same
	if cache.head != cache.tail {
		t.Error("Error removing the last entry of the cache")
	}
}

// func TestRemoveEntry_RemoveWhenCacheIsEmpty(t *testing.T) {}

// Others
