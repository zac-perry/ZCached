package main

import (
	"fmt"
	"zcached/internal/telnet"
)

func main() {
	fmt.Println("Hello, world!")

	server := telnet.NewServer("11211")
	server.Start()

	// start the telnet server
	// cmd line arguments -- port number (default is 11211)

	// server will need to create a socket, and bind it to the address of the server
	// IP address: 127.0.0.1 and specified port (ip is just local host)

	// then, server needs to listen for requests, accept incoming request, and receive the data
}
