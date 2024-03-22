package client

import (
	"bufio"

	"strings"
)

func (c *Client) GetUsername() (string, error)  {
// Send the prompt to enter the username
if _, err := c.Conn.Write([]byte(NamePrompt)); err != nil {
	return "", err
}

// Read the username from the client
username, err := bufio.NewReader(c.Conn).ReadString('\n')
if err != nil {
	return "", err
}

// Trim any leading or trailing whitespace from the username
username = strings.TrimSpace(username)
return username, nil

}