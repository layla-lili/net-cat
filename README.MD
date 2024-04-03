
# NET-CAT
```bash
📦net-cat
 ┣ 📂cmd
 ┃ ┣ 📜history.txt
 ┃ ┣ 📜main.go
 ┃ ┗ 📜welcome.txt
 ┣ 📂internal
 ┃ ┣ 📂app
 ┃ ┃ ┣ 📂chat
 ┃ ┃ ┃ ┣ 📜group.go
 ┃ ┃ ┃ ┗ 📜message.go
 ┃ ┃ ┣ 📂client
 ┃ ┃ ┃ ┗ 📜client.go
 ┃ ┃ ┣ 📂server
 ┃ ┃ ┃ ┣ 📜server.go
 ┃ ┃ ┃ ┗ 📜utils.go
 ┃ ┃ ┣ 📂ui
 ┃ ┃ ┃ ┗ 📜ui.go
 ┃ ┃ ┣ 📂utils
 ┃ ┃ ┃ ┗ 📜utils.go
 ┃ ┃ ┗ 📜app.go
 ┃ ┣ 📂interfaces
 ┃ ┃ ┗ 📜server_initializer.go
 ┃ ┣ 📂logging
 ┃ ┃ ┗ 📜logger.go
 ┃ ┗ 📂storage
 ┃ ┃ ┗ 📜storage.go
 ┣ 📜README.MD
 ┣ 📜client_script.bash.txt
 ┣ 📜go.mod
 ┗ 📜learning.md
```


 ## TODO
+ unit testing - timing for client connection before name prompt and no message send for so long
+ Can the Clients change their names?

+ Is the chat group informed if a Client changes his name?

+ Does the server produce logs about Clients activities?

+ Are the server logs saved into a file?

+ Is there more NetCat flags implemented?

+ Does the project present a Terminal UI using JUST this package : https://github.com/jroimartin/gocui?


