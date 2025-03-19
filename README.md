# ZCached
ZCached is a high-performance, distributed memory object caching system (its just a memcached clone)

### Goal
Goal for this project is to build my own memcached server. Intended to be used for  speeding up web applications by reducing the DB load.
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
    - Struct for each entry in the cache
        - make sure to handle concurrent updates (mutexes, locks, whole shabang)
    - the cache itself 
    - figure out how to go ahead and get this initialized
        - Want to support LRU, only a specific number of entries or something, etc. (need to figure out memory) 

- [x] support expire time with LRU cache 
- [x] fix shared cache bug (clients using their own instance) 
- [x] Finish GET
    - [ ] Update returns
- [x] Finish PUT
    - [x] Update returns
    - [ ] finish tests
- [ ] Clean up the client message handling / logging 
- [ ] clean up function args 
- [ ] TCP graceful shutdown 
- [ ] UPDATE, other functions
- [x] concurrency control

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
5. data block?

### Notes: 
Currently implementing both LRU and TTL for the cache. This is kind of redundant in most cases. This is purely for learning purposes. If you were to actually use this, the purpose of the LRU would be to maintain certain memory limits, while the TTL would ensure data freshness. In reality, I have no use case for this lol but just thought it would be fun. 
So far: 
    - get works
    - push front and remove entry works
    - expired and lazy removing works


### Future ideas
- [ ] Make the cache an actual package that can be used, k/v in mem
- [ ] Make ux not horrible, defaults for expiration times, etc.
- [ ] support storage of anything (pointer to struct, etc)
- [ ] downtime recovery (similar to gocache) -> reload previously used caches, etc.
- [ ] add backup clean up process using the time.Ticker lib
- [x] Look into if using LRU for an actual cache, with this sort of expiration stuff is alright. LRU really only required if I want to set a memory limit or cache entry limit (which doing this in terms of system stats may be better)

