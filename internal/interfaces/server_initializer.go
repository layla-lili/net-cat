// Package interfaces defines the interface for initializing and running the NetCat server.
package interfaces

import (
	"netcat/internal/app/client"
	"sync"
)

var (
	Clients []client.Client

	Mutex       sync.Mutex
	HistoryFile = "history.txt"

	NamePrompt     = "\n[ENTER YOUR NAME]: "
	WelcomeMessage string
)



// ServerInitializer represents the interface for initializing and running the server.
type ServerInitializer interface {
    // InitializeServer initializes the server with the specified address.
    InitializeServer(addr string) error

    // ListenAndServe starts the server and listens for incoming connections.
    ListenAndServe() error



}

