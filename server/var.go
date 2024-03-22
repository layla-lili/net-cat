package server

import "netcat/client"

var (
 Clients []client.Client
 NamePrompt = "\n[ENTER YOUR NAME]: "
 WelcomeMessage string
)


