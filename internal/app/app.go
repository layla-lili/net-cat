// Package app provides the main application logic for starting the NetCat server.
package app

import (
	"fmt"
	"log"
	"netcat/internal/app/server"
	"netcat/internal/app/utils"
	"netcat/internal/interfaces"
	"os"
)

// App represents the main application struct, which encapsulates the server initializer.
type App struct {
	serverInitializer interfaces.ServerInitializer
}

// NewApp creates a new instance of the application with the given server initializer.
func NewApp(serverInitializer interfaces.ServerInitializer) *App {
	return &App{
		serverInitializer: serverInitializer,
	}
}

// StartServer starts the NetCat server with the specified port.
func (a *App) StartServer(port int) {
	// Construct the address string using the specified port.
	addr := fmt.Sprintf(":%d", port)

	// Initialize the server with the constructed address.
	if err := a.serverInitializer.InitializeServer(addr); err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}

	// Start listening for incoming connections and serving them.
	if err := a.serverInitializer.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// RunServer is a convenience function to start the NetCat server using command-line arguments.
func RunServer() {
	// Parse the port number from command-line arguments.
	port := utils.ParsePortFromArgs(os.Args)

	// Create a new server initializer instance.
	serverInitializer := server.NewServer()

	// Create a new application instance with the server initializer.
	app := NewApp(serverInitializer)

	// Start the server with the specified port.
	app.StartServer(port)
}
