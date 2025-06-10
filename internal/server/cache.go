package server

import (
	"errors"
	"log"
	"sync"
	"time"
)

// Cache struct
// Maintains a map of entries for easy look up
// Maintain a doubly linked list for LRU eviction + cache size limiting
type Cache struct {
	entries  map[string]*Entry
	capacity int

	mu   sync.Mutex
	head *Entry
	tail *Entry
}

// TODO: allow a value to be "anything" -- int, string, list, etc.
// Can allow JSON, whatever other format potentially
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

func NewCache(capacity int) *Cache {
	return &Cache{
		entries:  make(map[string]*Entry),
		capacity: capacity,
		mu:       sync.Mutex{},
		head:     nil,
		tail:     nil,
	}
}

/*
Get the cache entry with the specified key
If the cache is not empty and it exists, return the value
also need to check expire time
*/
// todo: refactor returns
// Return the entire entry?
func (this *Cache) Get(key string) (int, error) {
	this.mu.Lock()
	defer this.mu.Unlock()

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

/* Set */
// TODO: add other args here
func (this *Cache) Set(key string, value int) (string, error) {
	this.mu.Lock()
	defer this.mu.Unlock()

	log.Print("Set() -- Calling put on key val: ", key, " ", value)
	log.Print("Set() -- Checking if key already exists")

	// if it exists, move to the front of the dll and return
	if entry, ok := this.entries[key]; ok {
		log.Print("Set() -- Key already exists")
		log.Print(entry)
		this.MoveEntryToFront(entry)
		return "EXISTS", errors.New("Cache entry already exists")
	}

	// initialize new entry
	// if cache is full - remove whatever is at the end.
	log.Print("Set() -- Checking if the cache is full")
	if len(this.entries) == this.capacity {
		log.Print("Cache is full. Removing the least recently used item")
		this.RemoveEntry(this.tail)
	}

	log.Print("Set() -- Making new entry")
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

/* Update(): update the value for the given key */
func (this *Cache) Update(key string, value int) (string, error) {
	// find the cache entry
	// Update the value
	// move to front, refresh time, etc.
	// can probably just call get here (assuming I update the returns) and then change, set.
	return "", nil
}

func (this *Cache) MoveEntryToFront(entry *Entry) error {
	log.Print("MoveEntryToFront() -- Moving entry to the front of the cache -- ", entry.key)
	if this.head == nil {
		this.head = entry
		this.tail = entry
		return nil
	} else if entry == this.head {
		return nil
	}

	if entry.Prev != nil {
		entry.Prev.Next = entry.Next
	}

	if entry.Next != nil {
		entry.Next.Prev = entry.Prev
	}

	if entry == this.tail {
		this.tail = entry.Prev
	}

	entry.Next = this.head
	entry.Prev = nil
	this.head.Prev = entry
	this.head = entry

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
	var entryNum int
	entry = this.head
	if entry == nil {
		log.Print("printList() -- Cache is currently empty.")
		return
	}

	entryNum = 1
	for entry != nil {
		log.Print(entryNum, ". Curr Key: ", entry.key)
		log.Print("-  Curr Next: ", entry.Next)
		log.Print("-  Curr Prev: ", entry.Prev)
		log.Print("------------------------------------")
		if entry.Next == entry || entry.Next == this.head {
			break
		}
		entry = entry.Next
		entryNum += 1
	}

	log.Print("Current head: ", this.head)
	log.Print("Current tail: ", this.tail)
}
