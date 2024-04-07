package client

import (
	"bufio"
	"net"
	"fmt"
	"os"
)

// Client represents a connected client
type Client struct {
	Name string
	Conn net.Conn
	// Messages chan string
	Writer *bufio.Writer
}

//Not Used
func ClientConnect() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:9898")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Create a bufio Reader to read messages from the server
	reader := bufio.NewReader(conn)

	// Create a bufio Writer to send messages to the server
	writer := bufio.NewWriter(conn)

	// Start a goroutine to continuously read messages from the server
	go func() {
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading:", err)
				return
			}
			fmt.Print("Server:", message)
		}
	}()

	// Start a loop to send messages to the server
	for {
		// Read user input from the console
		fmt.Print("Enter message: ")
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from console:", err)
			continue
		}

		// Send the message to the server
		_, err = writer.WriteString(message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			continue
		}
		writer.Flush() // Flush the buffer to ensure the message is sent immediately
	}
}
