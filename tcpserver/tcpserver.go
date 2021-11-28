package tcpserver

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	var m []string

	for {
		//message, err := reader.ReadString('\r')

		//var hexStartMessage byte = 0x0B

		var hexFinishMessage1 byte = 0x1C
		//var hexFinishMessage2 byte = 0x0D

		//var hexFinishSegment byte = 0x0D



		message, err := reader.ReadString(hexFinishMessage1)
		fmt.Printf("Message incoming: %s\n", message)
		//fmt.Println("Byte incoming: ", []byte(message)
		parseHL7(message)
		switch err {
		case nil:
			m = append(m, message)
		case io.EOF:
			log.Println("server closed the connection")
			fmt.Printf("Message after EOF : %s\n", m)
			client.conn.Write([]byte("Message received.\n"))
			client.conn.Close()
			return
		default:
			log.Printf("server error: %v\n", err)
			client.conn.Close()
			return
		}

		/*
			if err != nil {
				client.conn.Close()
				if err == io.EOF {
					fmt.Println("End of message reached")
				}
				return
			}
		*/

		// parseHL7(message)

		/*
			/// verschieben
			MSHmatched, _ := regexp.MatchString("^MSH", message)
			if MSHmatched {
				fmt.Println("MSH found: ", MSHmatched) // true or false
			} else {
				m = append(m, message)
			}
			client.conn.Write([]byte("Message received.\n"))
		*/
	}

}
