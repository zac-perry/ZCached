# ZCached
ZCached is a high-performance, distributed memory object caching system

### Goal
Goal for this project is to build my own memcached inspired server. Intended to be used for  speeding up web applications by reducing the DB load.
Memcached is an inmemory key-value store for small chunks of data retrieved from backend systems with higher latency. 


Memcached supports a client-server architecture where the client is aware of all servers, but the servers are not aware of one another.
IF a client wishes to set or read a value corresponding to a key, the client's libary computes a hash of the key to determine which server to use. Thisprovides a simple form of sharding and promotes a scalable, shared-nothing architecture across memcached servers.

When the server receives a request, it computes a second hash of the key to determine where to store/read the corresponding value. Values are stored in RAM and if it ever runs out, removes the oldest values. 

### How to run
- Compile and run the server
- connect using `nc localhost port`

### Current todos
- [x] Start parsing and accepting client input
- [x] Rename some folders/file stuff
- [x] Get the datastructure setup 
- [x] support expire time with LRU cache 
- [x] fix shared cache bug (clients using their own instance) 
- [x] Finish GET
    - [ ] Update returns
- [x] Finish PUT
    - [x] Update returns
    - [x] finish tests
- [ ] Clean up the client message handling / logging 
- [ ] clean up function args 
- [ ] TCP graceful shutdown 
- [ ] UPDATE, other functions
- [x] concurrency control
- [ ] Figure out expire time rotation (refresh time IF value is used?)

## MEMCACHE PROTOCOL
```
<command name> <key> <flags> <exptime> <byte count> [noreply]\r\n
<data block>\r\n
```

0. command name
1. key
2. flags
3. expTime
4. byte count
5. data block

### Notes: 
Currently implementing both LRU and TTL for the cache. This is kind of redundant in most cases. This is purely for learning purposes. If you were to actually use this, the purpose of the LRU would be to maintain certain memory limits, while the TTL would ensure data freshness. In reality, I have no use case for this lol but just thought it would be fun. 
So far: 
    - get works
    - push front and remove entry works
    - expired and lazy removing works


### Future ideas
- [ ] support storage of anything (pointer to struct, etc)
- [ ] downtime recovery (similar to gocache) -> reload previously used caches, etc.
- [ ] (maybe) add backup clean up process using the time.Ticker lib 



### Package ideas
- create this in the context of actually using it as a caching package
- Support different configs (LRU, some others maybe, TTL, etc) 
- Multiple opertaions
- Downtime recovery
- Concurrency support through mutexes, etc.

- I think this way, I can publish this as an actual go package. No one will probably use it (including me lol), but would be a good expeirence and nice to have for my portfolio

```Go
//Example of something I could .
// Cache represents a thread-safe caching system with TTL and LRU capabilities
type Cache[K comparable, V any] interface {
    // Basic operations
    Get(key K) (value V, found bool)
    Set(key K, value V) error
    SetWithTTL(key K, value V, ttl time.Duration) error
    Delete(key K) error
    Has(key K) bool
    
    // Batch operations
    GetMulti(keys []K) map[K]V
    SetMulti(items map[K]V) error
    DeleteMulti(keys []K) error
    
    // Cache management
    Clear() error
    Count() int
    Keys() []K
    
    // Advanced options
    SetMaxSize(size int)
    SetDefaultTTL(ttl time.Duration)
    OnEvicted(callback func(key K, value V))
}
```
