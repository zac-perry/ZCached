package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	port  string
	cache *Cache
}

type Client struct {
	conn   net.Conn
	server *Server
	id     int
}

func NewServer(port string) *Server {
	fmt.Println("Using port: ", port)

	return &Server{
		port:  port,
		cache: NewCache(5),
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
			id:     id,
			conn:   conn,
			server: server,
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
	// TODO: default this somewhere maybe? (env file?)
	//cache := NewCache(3)

	// Server stuff handling client messages
	for {
		// may need to change \n depending on query
		message, err := reader.ReadString('\n')
		messageParts := strings.Split(message, "\\r\\n")
		commands := strings.Fields(messageParts[0])
		data := []string{}

		if len(messageParts) > 1 {
			data = strings.Fields(messageParts[1])
		}

		for _, c := range commands {
			fmt.Println("part of command: ", c)
		}
		for _, d := range data {
			fmt.Println("part of data: ", d)
		}

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

		// TODO: make this a switch to handle and call the correct things?
		if commands[0] == "GET" {
			val, err := client.server.cache.Get(commands[1])
			if err != nil {
				client.conn.Write([]byte(err.Error()))
			}
			client.conn.Write([]byte(fmt.Sprintf("\nGET CALLED: %d \n", val)))
		} else if commands[0] == "PUT" {
			val, _ := strconv.Atoi(commands[2])
			msg, err := client.server.cache.Put(commands[1], val)
			if err != nil {
				client.conn.Write([]byte(err.Error()))
			}
			client.conn.Write([]byte(fmt.Sprintf("PUT CALLED: %s \n", msg)))
		}
	}
}
