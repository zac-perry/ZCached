package server

import (
	"errors"
	"log"
	"time"
)

// Want to support this format of messages:
// <command name> <key> <flags> <exptime> <byte count> [noreply]\r\n <data block>\r\n
// When a client sends the following, parse out the info and call the respective function?
// MAKE SURE TO HANDLE CACHE UPDATES AND ACCESSES CONCURRENTLY

// Cache struct
// Maintains a map of entries for easy look up
// Maintain a doubly linked list for LRU eviction + cache size limiting
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
func NewCache(capacity int) *Cache {
	return &Cache{
		make(map[string]*Entry),
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
// todo: refactor returns
func (this *Cache) Get(key string) (int, error) {
	if len(this.entries) == 0 {
		log.Print("Get() -- Cache is empty..")
		return -1, errors.New("Cache is empty")
	}

	if entry, ok := this.entries[key]; ok {
		log.Print("Get() -- Cache Entry found")

		if entry.isExpired() {
			log.Print("Get() -- Entry was expired and removed. Sorry!")
			this.RemoveEntry(entry)
			return -1, errors.New("Entry was expired")
		}

		// push to front
		this.MoveEntryToFront(entry)
		return entry.value, nil
	}

	return -1, errors.New("Entry not found")
}

/* Put */
// TODO: add other args here
func (this *Cache) Put(key string, value int) (string, error) {
	log.Print("Put() -- Calling put on key val: ", key, " ", value)
	log.Print("Put() -- Checking if key already exists")

	// if it exists, move to the front of the dll and return
	if entry, ok := this.entries[key]; ok {
		log.Print("Put() -- Key already exists")
		log.Print(entry)
		this.MoveEntryToFront(entry)
		return "EXISTS", nil
	}

	// initialize new entry
	// if cache is full - remove whatever is at the end.
	log.Print("Put() -- Checking if the cache is full")
	if len(this.entries) == this.capacity {
		log.Print("Cache is full. Removing the least recently used item")
		this.RemoveEntry(this.tail)
	}

	log.Print("Put() -- Making new entry")
	newEntry := &Entry{
		key:        key,
		value:      value,
		expireTime: nil,
		createdAt:  time.Now(),
		Next:       nil,
		Prev:       nil,
	}

	this.entries[key] = newEntry
	this.MoveEntryToFront(newEntry)

	return "STORED", nil
}

func (this *Cache) MoveEntryToFront(entry *Entry) error {
	log.Print("MoveEntryToFront() -- Moving entry to the front of the cache -- ", entry.key)
	if this.head == nil {
		this.head = entry
		this.tail = entry
		return nil
	}

	currHead := this.head
	this.head = entry
	this.head.Next = currHead
	currHead.Prev = this.head

	log.Print("MoveEntryToFront() -- New Head of Cache -- ", this.head.key)
	return nil
}

/*
RemoveEntry() -- will remove an entry from both the cache and the linked list
*/
func (this *Cache) RemoveEntry(entry *Entry) {
	log.Print("RemoveEntry() -- Removing an entry from the cache -- ", entry.key)

	if entry.Prev != nil {
		entry.Prev.Next = entry.Next
	}
	if entry.Next != nil {
		entry.Next.Prev = entry.Prev
	}

	if entry == this.head {
		this.head = entry.Next
	}

	if entry == this.tail {
		this.tail = entry.Prev
	}

	delete(this.entries, entry.key)

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

/*
printList() -- helper function for printing the dll out
*/
func (this *Cache) printList() {
	var entry *Entry

	entry = this.head
	for entry != nil {
		log.Print("Curr Key: ", entry.key)
		log.Print("  Curr Next: ", entry.Next)
		log.Print("  Curr Prev: ", entry.Prev)
		entry = entry.Next
	}
}
