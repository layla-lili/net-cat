package client

import (
	"bufio"
	"net"
)

// Client represents a connected client
type Client struct {
	Name string
	Conn net.Conn
	// Messages chan string
	Writer *bufio.Writer
}

//TODO REFACTOR
