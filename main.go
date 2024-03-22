package main

import (
	"fmt"
	"log"
	"netcat/server"
	"os"
)

func main() {
	// Parse command-line arguments for port number
	port := server.ParsePortFromArgs(os.Args)

	// Create a new server instance
	server := server.NewServer(fmt.Sprintf(":%d", port))

	// Start the server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}


