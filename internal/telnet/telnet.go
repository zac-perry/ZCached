package telnet

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	fmt.Println("Using port: ", port)
	return &Server{
		port: port,
	}
}

func (server *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", server.port))
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// client stuff
		client := &Client{
			conn: conn,
		}

		go client.handleRequest()
	}
}
