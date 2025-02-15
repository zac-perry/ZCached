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
	capacity int

	head *Entry
	tail *Entry
}

type Entry struct {
	key        string
	value      int // remove eventually
	flags      uint16
	data       []byte
	expireTime *time.Duration
	createdAt  time.Time
	Next       *Entry
	Prev       *Entry
}

// TODO: simplify this
func NewCache(capacity int) Cache {
	return Cache{
		make(map[string]*Entry, 0),
		capacity,
		nil,
		nil,
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
		if entry.isExpired() {
			log.Print("Entry was expired and removed. Sorry! \n")
			// todo: remove the entry
			return -1
		}

		// push to front
		return entry.value
	}

	// TODO: send this message to the client..
	log.Print("entry not found")
	return -1
}

/* Put */
// TODO: add other args here
func (this *Cache) Put(key string, value int) {
	log.Print("Put function not implemented\n")

	// make sure it doesn't exist already
	// if so, update the record and refresh the TTL
	// push to the front of the list

	// otherwise, insert, set fields
	// if the cache is full, remove the least recently used (pop back)
	// make sure to set the TTL

	if entry, ok := this.entries[key]; ok {
		log.Print("PUT -- Key already exists")
		log.Print(entry)
		// update the entry
		// set everthing
		// reset TTL, etc
		// push to the front
		return
	}

	// initialize new entry
	// if cache is full - remove whatever is at the end.
	if len(this.entries) == this.capacity {
		log.Print("Cache is full. Removing the least recently used item")
	}

	// insert, push front, return
	// when inserting, make sure to handle the head and tail here
}

// TODO
func (this *Cache) MoveEntryToFront(entry *Entry) error {
	return nil
}

/*
RemoveLRUEntry() will remove the least recently used (LRU) entry from the cache.
Does this by removing the tail of the linked list and deleting the entry from the map
*/
func (this *Cache) RemoveLRUEntry() {
	log.Print("RemoveEntry() -- Removing LRU entry in the cache -- ", this.tail.key)

	if this.tail == nil {
		return
	}

	entryKey := this.tail.key
	this.tail = this.tail.Prev
	if this.tail == nil {
		this.head = nil
	} else {
		this.tail.Next = nil
	}

	delete(this.entries, entryKey)

	log.Print("RemoveEntry() -- Cache after removing LRU: ", this.entries)
	return
}

/*
isExpired(), returns if the current entry is expired and needs to be removed..
*/
func (this *Entry) isExpired() bool {
	if this.expireTime == nil {
		return false
	}

	elapsedTime := time.Since(this.createdAt)
	return elapsedTime > *this.expireTime
}

// TODO: handle when to remove an expired entry
