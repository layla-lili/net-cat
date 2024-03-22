package client

import "net"

// Client represents a connected client
type Client struct {
    Name     string
    Conn     net.Conn
    Messages chan string
}

// Connection represents a connection to the server
type Connection struct {
    Port string
    Host string
}
