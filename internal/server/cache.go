package server

import (
	"log"
	"time"
)

// Want to support this format of messages:
// <command name> <key> <flags> <exptime> <byte count> [noreply]\r\n <data block>\r\n
// When a client sends the following, parse out the info and call the respective function?
// MAKE SURE TO HANDLE CACHE UPDATES AND ACCESSES CONCURRENTLY

// Cache struct
// Mainaints a map of entries for easy look up
// Cache itself is a doubly linked list
type Cache struct {
	entries  map[string]*Entry
	sentinel *Entry
	capacity int
}

type Entry struct {
	key        string
	value      int // remove eventually
	flags      uint16
	data       []byte
	expireTime int64
	Next       *Entry
	Prev       *Entry
}

func Constructor(capacity int) Cache {
	return Cache{
		make(map[string]*Entry, 0),
		&Entry{Next: nil, Prev: nil},
		capacity,
	}
}

/*
Get the cache entry with the specified key
If the cache is not empty and it exists, return the value
also need to check expire time
*/
func (this *Cache) Get(key string) int {
	if len(this.entries) == 0 {
		log.Print("Cache is empty..\n")
		return -1
	}

	if entry, ok := this.entries[key]; ok {
		log.Print("Value found\n")
		if entry.Expired() {
			log.Print("Entry was expired and removed. Sorry! \n")
			// todo: remove the entry
			return -1
		}
		return entry.value
	}

	log.Print("Get function not implemented\n")
	return -1
}

/* Put */
func (this *Cache) Put(key string, value int) {
	log.Print("Put function not implemented\n")
}

/* Expired, returns if the current entry is expired and needs to be removed.. */
func (this *Entry) Expired() bool {

	if this.expireTime == 0 {
		return false
	}

	return time.Now().Unix() > this.expireTime
}

// Push front
// Pop
