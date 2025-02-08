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

### Current todos (as of 1/16) 
- [x] Start parsing and accepting client input
- [x] Rename some folders/file stuff
- [ ] Get the datastructure setup 
    - Struct for each entry in the cache
        - make sure to handle concurrent updates (mutexes, locks, whole shabang)
    - the cache itself 
    - figure out how to go ahead and get this initialized
        - Want to support LRU, only a specific number of entries or something, etc. (need to figure out memory) 

- [ ] support expire time with LRU cache 
    

## MEMCACHE PROTOCOL
```
<command name> <key> <flags> <exptime> <byte count> [noreply]\r\n
<data block>\r\n
```
