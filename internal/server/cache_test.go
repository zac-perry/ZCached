package server

import (
	"testing"
	"time"
)

// Get

// Put

/* Expire Tests */
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

// Pop

// Others
