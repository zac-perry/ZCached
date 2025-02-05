package server

import (
	"log"
)

// Want to support this format of messages:
// <command name> <key> <flags> <exptime> <byte count> [noreply]\r\n <data block>\r\n
// When a client sends the following, parse out the info and call the respective function?
// MAKE SURE TO HANDLE CACHE UPDATES AND ACCESSES CONCURRENTLY

// Cache struct
// Mainaints a map of entries for easy look up
// Cache itself is a doubly linked list
type Cache struct {
	entries  map[int]*Entry
	sentinel *Entry
	capacity int
}

// TODO: Update to include needed things for memcache request
type Entry struct {
	key   int
	flags []byte
	data  []byte
	Next  *Entry
	Prev  *Entry
}

func Constructor(capacity int) Cache {
	return Cache{
		make(map[int]*Entry, 0),
		&Entry{Next: nil, Prev: nil},
		capacity,
	}
}

// Get
func (this *Cache) Get(key int) int {
	log.Print("Get function not implemented\n")
	return -1
}

// Put
func (this *Cache) Put(key int, value int) {
	log.Fatalf("Put function not implemented\n")
}

// Push front

// Pop
