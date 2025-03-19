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

	// TODO: determine where to set the cache size default i guess
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

	// TODO:
	// If i want to do a graceful shutdown type thing, will have to add WG here to add client processes.
	// Whenever shutting down, will need to make sure to clear out client connections that still exist, etc.
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
			fmt.Println("DEBUG -- part of command: ", c)
		}
		for _, d := range data {
			fmt.Println("DEBUG -- part of data: ", d)
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

		/*	fmt.Println("Messaged Recieved from client #", client.id, ": ", message)
			client.conn.Write([]byte("Message recieved!\n"))*/

		switch commands[0] {
		// TOOD: Add update, etc.
		case "GET":
			val, err := client.server.cache.Get(commands[1])
			if err != nil {
				client.conn.Write([]byte(err.Error()))
			}
			client.conn.Write([]byte(fmt.Sprintf("\nGET CALLED: %d \n", val)))

		case "SET":
			val, _ := strconv.Atoi(commands[2])
			msg, err := client.server.cache.Set(commands[1], val)
			if err != nil {
				client.conn.Write([]byte(err.Error()))
			}
			client.conn.Write([]byte(fmt.Sprintf("GET CALLED: %s \n", msg)))

		case "UPDATE":
			client.conn.Write([]byte(fmt.Sprintf("UPDATE CALLED: (NOT IMPLEMENTED)\n")))

		case "PRINT":
			client.server.cache.printList()

		default:
			client.conn.Write([]byte(fmt.Sprintf("\n Command not specified or supported..\n")))
		}
	}
}
