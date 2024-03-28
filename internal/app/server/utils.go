package server

import (
	"fmt"
	"io/ioutil"
	"netcat/internal/interfaces"
	"os"
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
		return nil, fmt.Errorf("error reading history file: %v", err)
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
		return fmt.Errorf("error opening history file: %v", err)
	}
	defer file.Close()

	// Write the message to the file
	_, err = file.WriteString(message)
	if err != nil {
		return fmt.Errorf("error writing to history file: %v", err)
	}

	return nil
}