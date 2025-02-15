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
func TestGet_ValueExists(t *testing.T) {
	cacheEntry1 := Entry{
		key:   "1",
		value: 100,
	}

	cacheEntry2 := Entry{
		key:   "2",
		value: 200,
	}

	cacheEntry3 := Entry{
		key:   "3",
		value: 300,
	}
	cache := NewCache(3)
	cache.head = &cacheEntry1
	cache.tail = &cacheEntry2
	cache.head.Next = cache.tail
	cache.tail.Prev = cache.head
	cache.entries[cacheEntry1.key] = &cacheEntry1
	cache.entries[cacheEntry2.key] = &cacheEntry2
	cache.entries[cacheEntry3.key] = &cacheEntry3

	val, err := cache.Get("2")

	if val != 200 || err != nil {
		t.Error("Error getting the cache entry")
	}
}

func TestGet_ValueDoesNotExist(t *testing.T) {
	cacheEntry1 := Entry{
		key:   "1",
		value: 100,
	}

	cacheEntry2 := Entry{
		key:   "2",
		value: 200,
	}

	cacheEntry3 := Entry{
		key:   "3",
		value: 300,
	}
	cache := NewCache(3)
	cache.head = &cacheEntry1
	cache.tail = &cacheEntry2
	cache.head.Next = cache.tail
	cache.tail.Prev = cache.head
	cache.entries[cacheEntry1.key] = &cacheEntry1
	cache.entries[cacheEntry2.key] = &cacheEntry2
	cache.entries[cacheEntry3.key] = &cacheEntry3

	_, err := cache.Get("4")

	if err == nil {
		t.Error("Value does not exist check failed")
	}
}

func TestGet_CacheIsEmpty(t *testing.T) {
	cache := NewCache(3)

	_, err := cache.Get("4")

	if err == nil {
		t.Error("Value does not exist check failed")
	}
}

func TestGet_ValueIsExpired(t *testing.T) {
	expire := time.Second
	cacheEntry1 := Entry{
		key:        "1",
		value:      100,
		createdAt:  time.Now().Add(-2 * time.Second),
		expireTime: &expire,
	}
	cache := NewCache(3)
	cache.entries[cacheEntry1.key] = &cacheEntry1

	_, err := cache.Get("1")

	if err == nil {
		t.Error("Value does not exist check failed")
	}
}

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

# MoveEntryToFront Function Tests

--------------------------
*/
func TestMoveEntryToFront_MoveEntry(t *testing.T) {
	cacheEntry1 := Entry{
		key:   "1",
		value: 100,
	}

	cacheEntry2 := Entry{
		key:   "2",
		value: 200,
	}

	cacheEntry3 := Entry{
		key:   "3",
		value: 300,
	}
	cache := NewCache(3)
	cache.head = &cacheEntry1
	cache.tail = &cacheEntry2
	cache.head.Next = cache.tail
	cache.tail.Prev = cache.head
	cache.entries[cacheEntry1.key] = &cacheEntry1
	cache.entries[cacheEntry2.key] = &cacheEntry2
	cache.entries[cacheEntry3.key] = &cacheEntry3

	cache.MoveEntryToFront(&cacheEntry3)

	// ensure the head and stuff is correct.
	if cache.head.key != cacheEntry3.key ||
		cache.head.Next.key != cacheEntry1.key {
		t.Error("Error moving entry to the front of the linked list")
	}
}

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
	cache.entries[cacheEntry1.key] = &cacheEntry1
	cache.entries[cacheEntry2.key] = &cacheEntry2

	cache.RemoveEntry(cache.tail)

	if len(cache.entries) != 1 {
		t.Error("Error removing the last entry of the cache")
	}

	// in this case, head and tail should be the same
	if cache.head != cache.tail {
		t.Error("Error removing the last entry of the cache")
	}
}

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

/*
--------------------------

# Random Helper Function Tests

--------------------------
*/
func TestPrintingList_RandomTestToCheckLL(t *testing.T) {
	cacheEntry1 := Entry{
		key:   "1",
		value: 100,
	}

	cacheEntry2 := Entry{
		key:   "2",
		value: 200,
	}

	cacheEntry3 := Entry{
		key:   "3",
		value: 300,
	}
	cache := NewCache(3)
	cache.head = &cacheEntry1
	cache.tail = &cacheEntry2
	cache.head.Next = cache.tail
	cache.tail.Prev = cache.head
	cache.entries[cacheEntry1.key] = &cacheEntry1
	cache.entries[cacheEntry2.key] = &cacheEntry2
	cache.entries[cacheEntry3.key] = &cacheEntry3

	cache.MoveEntryToFront(&cacheEntry3)
	cache.printList()
}
