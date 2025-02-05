package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Server struct {
	port string
}

type Client struct {
	conn net.Conn
	id   int
}

func NewServer(port string) *Server {
	fmt.Println("Using port: ", port)
	return &Server{
		port: port,
	}
}

/*
Start up the TCP server. Accept client connections
*/
func (server *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", server.port))
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()
	id := 1

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// client stuff
		client := &Client{
			id:   id,
			conn: conn,
		}

		go client.handleRequest()
		id++
	}
}

/*
Handle incoming client messages.
*/
func (client *Client) handleRequest() {
	fmt.Println("Client #", client.id, " connected..")
	client.conn.Write([]byte("Connected!\n"))

	reader := bufio.NewReader(client.conn)
	// TODO: default this somewhere maybe?
	cache := Constructor(5)

	// Server stuff handling client messages
	for {
		// may need to change \n depending on query
		message, err := reader.ReadString('\n')
		messageParts := strings.Split(strings.TrimSpace(message), " ")

		if err != nil {
			if err == io.EOF {
				fmt.Println("Client #", client.id, " disconnected..")
				client.conn.Close()
				return

			}
			client.conn.Write([]byte("Error encountered...\n"))
			client.conn.Close()
			return
		}

		fmt.Println("Messaged Recieved from client #", client.id, ": ", message)
		client.conn.Write([]byte("Message recieved!\n"))

		if messageParts[0] == "GET" {
			val := cache.Get("1")
			client.conn.Write([]byte(fmt.Sprintf("GET CALLED: %d \n", val)))
		}

	}
}
