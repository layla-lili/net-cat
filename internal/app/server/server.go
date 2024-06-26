// Package server provides functionality for managing the NetCat server,
// handling client connections, and facilitating communication between clients.
package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"netcat/internal/app/client"
	"netcat/internal/interfaces"
	"netcat/internal/logging"
)

// Server represents a TCP server for the NetCat application.
type Server struct {
	Addr             string     // Address on which the server listens for incoming connections
	Mutex            sync.Mutex // Mutex for thread-safe access to server state
	ActiveClients    int        // Number of active clients connected to the server
	ActiveClientsMux sync.Mutex // Mutex for thread-safe access to active client count
}

// init initializes the server by reading welcome message and setting history file.
func init() {
	// Read the content of welcome.txt
	content, err := os.ReadFile("welcome.txt")
	if err != nil {
		logging.Logger(err.Error())
		log.Fatalf("Error reading welcome.txt: %v", err)
	} else {
		logging.Logger("Welcome file read successfully")
	}
	interfaces.WelcomeMessage = string(content)

	interfaces.HistoryFile = "history.txt"
}

// NewServer creates a new instance of the server.
func NewServer() interfaces.ServerInitializer {
	return &Server{}
}

// InitializeServer initializes the server with the specified address.
func (s *Server) InitializeServer(addr string) error {
	if s == nil {
		log.Println("Error: nil")
		return nil
	}

	s.CleanupHistoryFile()
	s.Addr = addr
	return nil
}

// ListenAndServe starts the server and listens for incoming connections.
func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)

	if err != nil {
		logging.Logger(err.Error())
		return fmt.Errorf("failed to listen: %v", err)
	} else {
		logging.Logger("Server listening on " + s.Addr)
	}

	defer listener.Close()

	// log.Printf("Server listening on %s", s.Addr)
	log.Printf("Listening on the IP: %s and port :%d", GetIpLocal(), ParsePortFromArgs(os.Args))

	for {
		conn, err := listener.Accept()

		if err != nil {
			logging.Logger(err.Error())
			log.Printf("Error accepting connection: %v", err)
			continue
		} else {
			logging.Logger("Connection accepted successfully")
		}
		go s.handleConnection(conn)
	}
}

// handleConnection handles an incoming connection from a client.
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("New connection from %s", conn.RemoteAddr())

	s.sendWelcomeMessage(conn)
	username := s.promptUsername(conn)

	// Attempt to add the client
	err1 := s.addClient(conn, username)
	if err1 != nil {
		log.Printf("Error adding client: %v", err1)
		conn.Write([]byte("Sorry, the chat room is full. Please try again later.\n"))
		return
	}

	// log.Printf("Client '%s' added", username)

	// Send join message to all clients except the new client
	joinMessage := fmt.Sprintf("\n%s has joined our chat...\n", username)
	s.broadcast(joinMessage, "", conn, "join")

	// Load history messages for the newly joined client
	s.loadHistoryMessages(conn)

	// Send initial message template only to the new client
	s.sendInitialMessages(conn, username)

	// Handle client messages
	s.handleClientMessages(conn, username)

	// If client disconnects, remove it from the list and broadcast leave message
	err := s.removeClient(conn)
	if err != nil {
		logging.Logger(err.Error())
		log.Printf("Error removing client: %v", err)
	} else {
		logging.Logger("Client removed successfully")

	}

	leaveMessage := fmt.Sprintf("\n%s has left our chat...\n", username)
	s.broadcast(leaveMessage, "", conn, "leave")
}

// handleClientMessages handles messages received from a client.
func (s *Server) handleClientMessages(conn net.Conn, username string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		errMsg := verifyMessage(message)
		if errMsg != "" {
			// Send error message to the client
			conn.Write([]byte(errMsg + "\n"))
			s.sendReadyMessages(conn, username)
			continue
		}

		time := time.Now().Format("2006-01-02 15:04:05")
		// Broadcast the message to all clients except the sender
		formattedMessage := fmt.Sprintf("\n[%s][%s]: %s\n", time, username, message)
        // Send the message to all clients except the sender
		s.broadcast(formattedMessage+"hcm", username, conn, "formatted")
		// Send ready message to the client himself
		s.sendReadyMessages(conn, username)

		// Save the message to history
		if err := SaveHistoryMessage(formattedMessage + "\n"); err != nil {
			logging.Logger(err.Error())
			log.Printf("Error saving message to history: %v", err)
		} else {
			logging.Logger("Message saved to history successfully")

		}
	}

	if err := scanner.Err(); err != nil {
		logging.Logger(err.Error())
		log.Printf("Error reading from %s: %v", username, err)
	} else {
		logging.Logger("Error reading from client")

	}

	log.Printf("%s disconnected", username)
}

// sendWelcomeMessage sends the welcome message to a newly connected client.
func (s *Server) sendWelcomeMessage(conn net.Conn) {
	if s == nil {
		log.Println("Error: Mock connection is nil or CloseFunc is nil")
		return
	}
	conn.Write([]byte(interfaces.WelcomeMessage))
}

// addClient adds a new client to the server.
func (s *Server) addClient(conn net.Conn, username string) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	// Check if the maximum number of clients has been reached
	if len(interfaces.Clients) >= 10 { // Check for maximum limit
		log.Println("Client tried to connect, but no space available.")
		return fmt.Errorf("maximum client limit reached")
	}

	// Add the client to the list of active clients (inside the critical section)
	s.addClientToList(conn, username)
	log.Printf("Client '%s' added successfully", username)
	return nil
}

func (s *Server) addClientToList(conn net.Conn, username string) {
	writer := bufio.NewWriter(conn)
	interfaces.Clients = append(interfaces.Clients, client.Client{Conn: conn, Name: username, Writer: writer})

	// Increment the active client count (inside the critical section)
	s.ActiveClientsMux.Lock()
	s.ActiveClients++
	s.ActiveClientsMux.Unlock()
}

// sendInitialMessages sends initial messages to a newly connected client.
func (s *Server) sendInitialMessages(conn net.Conn, username string) {
	// Send the template message to the newly joined client
	templateMessage := fmt.Sprintf("\n[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), username)
	conn.Write([]byte(templateMessage))
}

// sendReadyMessages sends ready messages to a client himself.
func (s *Server) sendReadyMessages(conn net.Conn, username string) {
	templateMessage := fmt.Sprintf("\n[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), username)
	conn.Write([]byte(templateMessage))
}

// loadHistoryMessages loads history messages for a newly connected client.
func (s *Server) loadHistoryMessages(conn net.Conn) {
	historyMessages, err := LoadHistoryMessages()
	if err != nil {
		logging.Logger(err.Error())
		log.Printf("Error loading history messages: %v", err)
		return
	} else {
		logging.Logger("History messages loaded successfully")
	}
	writer := bufio.NewWriter(conn)
	for _, message := range historyMessages {
		writer.WriteString(message + "\n")
		writer.Flush()
	}
}

// broadcast sends a message to all connected clients except the sender.
func (s *Server) broadcast(message string, senderName string, sender net.Conn, messageType string) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	for _, client := range interfaces.Clients {
		if client.Conn != sender {
			switch messageType {
			case "formatted":
				//outgoing connection
				templateMessage := string([]byte(fmt.Sprintf("\n[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), client.Name)))

				// Broadcast formatted message to client
				if client.Conn != sender {
					client.Writer.WriteString(message + templateMessage)
				} else {
					client.Writer.WriteString(templateMessage)
				}
				client.Writer.Flush()

			case "join":
				templateMessage := string([]byte(fmt.Sprintf("\n[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), client.Name)))
				// Send join message to client
				if client.Conn != sender {
					client.Writer.WriteString(message + templateMessage)
				}
				client.Writer.Flush()

			case "leave":
				templateMessage := string([]byte(fmt.Sprintf("\n[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), client.Name)))

				// Send leave message to client
				if client.Conn != sender {
					client.Writer.WriteString(message + templateMessage)
				}
				client.Writer.Flush()

			case "ready":
				// Handle "ready" message differently
				if client.Conn != sender {
					// Display "ready" message to the client
					client.Writer.WriteString(message)
					client.Writer.Flush()
				} else {
					client.Writer.WriteString(message)
					client.Writer.Flush()
				}
			default:
				// Invalid message type
				log.Printf("Invalid message type: %s", messageType)
				return
			}
			// Release the mutex temporarily while formatting and writing messages
			s.Mutex.Unlock()
			client.Writer.Flush()
			s.Mutex.Lock()

		}
	}
}

// removeClient removes a disconnected client from the list of connected clients.
func (s *Server) removeClient(conn net.Conn) error {
	for i, client := range interfaces.Clients {
		if client.Conn == conn {
			// Remove the client from the list
			interfaces.Clients = append(interfaces.Clients[:i], interfaces.Clients[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("client not found")
}

// CleanupHistoryFile clears the history file.
func (s *Server) CleanupHistoryFile() error {
	// Open the history file in write mode
	file, err := os.OpenFile("history.txt", os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		logging.Logger(err.Error())
		return fmt.Errorf("failed to open history file: %v", err)
	} else {
		logging.Logger("History file opened successfully")
	}
	defer file.Close()

	// Truncate the file to 0 bytes, effectively clearing its contents
	if err := file.Truncate(0); err != nil {
		logging.Logger(err.Error())
		return fmt.Errorf("failed to truncate history file: %v", err)
	} else {
		logging.Logger("History file truncated successfully")

	}

	return nil
}

// promptUsername prompts the client to enter a username.
func (s *Server) promptUsername(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	for {
		conn.Write([]byte(interfaces.NamePrompt))
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		if err := isValidUsername(username); err != nil {
			logging.Logger(err.Error())
			// Invalid username, prompt again
			conn.Write([]byte(err.Error() + "\n"))
		} else {
			// Valid username, return it
			return username
			logging.Logger("Username entered successfully")
		}
	}
}
