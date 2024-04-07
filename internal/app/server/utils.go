package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"netcat/internal/interfaces"
	"netcat/internal/logging"
	"os"
	"strconv"
	"strings"
)

func verifyMessage(message string) string {
	if message == "" {
		return "you can't send empty messages"
	}
	return ""
}

// isValidUsername checks if the provided username is valid.
func isValidUsername(username string) error {
	// Check if the username is empty
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	// Check if the username length is more than 15 characters
	if len(username) > 15 {
		return fmt.Errorf("username cannot be more than 15 characters long")
	}

	// Check if the username already exists
	for _, client := range interfaces.Clients {
		if client.Name == username {
			return fmt.Errorf("username already exists")
		}
	}

	// If all checks pass, return nil indicating the username is valid
	return nil
}

// loadHistoryMessages reads previous messages from a history file.
func LoadHistoryMessages() ([]string, error) {
	// Read the content of the history file
	content, err := ioutil.ReadFile("history.txt")

	if err != nil {
		logging.Logger(err.Error())
		return nil, fmt.Errorf("error reading history file: %v", err)
	}else{
		logging.Logger("History file read successfully")
	
	}

	// Split the content into individual messages
	messages := strings.Split(string(content), "\n")

	return messages, nil
}

// saveHistoryMessage appends a message to the history file.
func SaveHistoryMessage(message string) error {
	// Open the history file in append mode
	file, err := os.OpenFile("history.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		logging.Logger(err.Error())
		return fmt.Errorf("error opening history file: %v", err)
	}else{
		logging.Logger("History file opened successfully")
	
	}
	defer file.Close()

	// Write the message to the file
	_, err = file.WriteString(message)

	if err != nil {
		logging.Logger(err.Error())
		return fmt.Errorf("error writing to history file: %v", err)
	}else{
		logging.Logger("History file written successfully")
	}

	return nil
}

func GetIpLocal() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	
	if err != nil {
		logging.Logger(err.Error())
		log.Fatal(err)
	}else{
		logging.Logger("UDP connection established successfully")
	
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP.String()
}
// parsePortFromArgs parses the port number from command-line arguments.
// If no port number is provided, it defaults to 8989.
func ParsePortFromArgs(args []string) int {
	if len(args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(1)
	}

	if len(args) == 1 {
		return 8989
	}

	port, err := strconv.Atoi(args[1])
	
	if err != nil || port < 1 || port > 65535 {
		logging.Logger(err.Error())
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(1)
	}else{
		logging.Logger("Port number parsed successfully")
	}

	return port
}