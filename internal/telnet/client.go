package telnet

import (
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
}

func (client *Client) handleRequest() {

	fmt.Println("Client connected?")
	client.conn.Write([]byte("Trying....\n"))
	client.conn.Write([]byte("Connected?\n"))
}
