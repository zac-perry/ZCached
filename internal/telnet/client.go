package telnet

import "net"

type Client struct {
	conn net.Conn
}

func (client *Client) handleRequest() {

	// read from buffer
	client.conn.Write([]byte("Connected?\n"))
}
