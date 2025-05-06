package main

import (
	"flag"
	"strconv"
	"zcached/internal/server"
)

func main() {
	// start the telnet server
	// cmd line arguments -- port number (default is 11211)

	// server will need to create a socket, and bind it to the address of the server
	// IP address: 127.0.0.1 and specified port (ip is just local host)

	// then, server needs to listen for requests, accept incoming request, and receive the data

	port := flag.Int("port", 11211, "port number")
	flag.Parse()

	// TODO: maybe set cache size here?
	server := server.NewServer(strconv.Itoa(*port))
	server.Start()
}
