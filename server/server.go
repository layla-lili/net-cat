package server

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
	"netcat/client"
)



func init() {
    // Read the content of welcome.txt
    content, err := ioutil.ReadFile("welcome.txt")
    if err != nil {
        log.Fatalf("Error reading welcome.txt: %v", err)
    }
    WelcomeMessage = string(content)
}


// NewServer creates a new TCP server with the specified address.
func NewServer(addr string) *Server {
	return &Server{Addr: addr}
}

// ListenAndServe starts the TCP server and handles incoming connections.
func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	defer listener.Close()

	log.Printf("Server listening on %s", s.Addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

// handleConnection handles an incoming connection.
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("New connection from %s", conn.RemoteAddr())
     //When a new client connects to the server, add it to the list of connected clients
	 conn.Write([]byte(WelcomeMessage))
	//  conn.Write([]byte(NamePrompt))

	 client := &client.Client{
		Conn:     conn,
		Messages: make(chan string),
	}
	
    // Obtain username from the client
    username, _ := client.GetUsername()

 // Format the message with current time and username
 formattedMessage := fmt.Sprintf("[%s][%s]: %s", time.Now().Format("2006-01-02 15:04:05"), username, "")
 
for {
    conn.Write([]byte(formattedMessage))
    message, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        // Handle error
        break
    }
    // Format the message with current time and username
    formattedMessage := fmt.Sprintf("[%s][%s]: %s", time.Now().Format("2006-01-02 15:04:05"), username, strings.TrimSpace(message))
    // Broadcast the message to other clients
    broadcast(formattedMessage, conn)
}


}



