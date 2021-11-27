package tcpserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
)

// Server ...
type Server struct {
	host string
	port string
}

// Client ...
type Client struct {
	conn net.Conn
}

// Config ...
type Config struct {
	Host string
	Port string
}

// New ...
func NewServer(config *Config) *Server {
	return &Server{
		host: config.Host,
		port: config.Port,
	}
}

// Run ...
func (server *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("HL7 TCP Server start listening...")
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{
			conn: conn,
		}
		go client.handleRequest()
	}
}

func (client *Client) handleRequest() {
	reader := bufio.NewReader(client.conn)
	fmt.Println("Got connection: ", client.conn.RemoteAddr())
	for {
		message, err := reader.ReadString('\n')
		//message, err := reader.ReadString('|')
		if err != nil {
			client.conn.Close()
			return
		}
		fmt.Printf("Message incoming: %s\n", string(message))
		parseHL7(message)
		client.conn.Write([]byte("Message received.\n"))
	}
}

func parseHL7(message string) {
	fmt.Println("HL7 parser starting...")
	//s := strings.Split(message, "|")
	s := regexp.MustCompile(`MSH|PID|^`)
	fmt.Printf("%q\n", s.Split(message, -1))
}
