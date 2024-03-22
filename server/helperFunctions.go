package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

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
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(1)
	}

	return port
}

func broadcast(message string, sender net.Conn) {
    for _, client := range Clients{
        if client.Conn != sender {
            client.Conn.Write([]byte(message))
        }
    }
}





