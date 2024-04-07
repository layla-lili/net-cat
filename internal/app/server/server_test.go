package server

import (
	"fmt"
	"net"
	colortest "netcat/internal/app/colorTest"
	mocks "netcat/internal/app/mocks"
	"netcat/internal/interfaces"
	"strings"
	"testing"
	"time"
)

// TestNewServer tests the NewServer function.
func TestNewServer(t *testing.T) {
	// Run subtests using t.Run
	colortest.LogInfo(t, "Running TestNewServer...")
	// Test NewServer function
	srv := NewServer()
	// Assert that srv is of type *server.Server
	_, ok := srv.(*Server)
	if !ok {
		colortest.LogError(t, "NewServer() did not return an instance of *server.Server")
	} else {
		colortest.LogSuccess(t, "TestNewServer complete successfully")
	}
}

// TestInitializeServer tests the InitializeServer function.
func TestInitializeServer(t *testing.T) {
	colortest.LogInfo(t, "Running TestInitializeServer...")
	// Create a new instance of Server
	srv := NewServer().(*Server) // Cast to Server type

	// Call InitializeServer with a test address
	err := srv.InitializeServer("127.0.0.1:8989")
	if err != nil {
		t.Errorf( "\033[31m"+"InitializeServer failed %v"+ "\033[0m", err)
	}
	// Assert that the server's address is set correctly
	expctedAddr := "127.0.0.1:8989"
	if srv.Addr != expctedAddr {
		colortest.LogError(t, "expcted server address "+expctedAddr+" but got "+srv.Addr)
		return // Exit the test early since it failed
	}
	// If the test reaches this point without any errors, it has passed
	colortest.LogSuccess(t, "InitializeServer completed successfully")

}

func TestSendWelcomeMessage(t *testing.T) {
    colortest.LogInfo(t, "Running TestSendWelcomeMessage...")
    // Create a new Server instance
    srv := NewServer().(*Server)

    // Create a mock net.Conn object
    mockConn := &mocks.MockConn{
    WriteFunc: func(p []byte) (n int, err error) {
        // Verify that the welcome message is sent
        expectedWelcomeMessage := interfaces.WelcomeMessage
        if string(p) != expectedWelcomeMessage {
            colortest.LogError(t,"Expected welcome message: "+expectedWelcomeMessage+" got: "+string(p))
        } else {
            colortest.LogSuccess(t,"TestSendWelcomeMessage completed sucessfully")
        }
        return len(p), nil
    },
}
    // Call sendWelcomeMessage with the mock connection
    srv.sendWelcomeMessage(mockConn)
}

//TestPromptUsername tests the promptUsername function.
func TestPromptUsername(t *testing.T) {
	colortest.LogInfo(t, "Running TestPromptUsername...")
    // Create a mock connection
    mockConn := &mocks.MockConn{
        Reader: strings.NewReader(""), // Pass an empty string reader
        WriteFunc: func(p []byte) (n int, err error) {
            // Verify the output prompt message
            expectedPrompt := interfaces.NamePrompt
            if string(p) != expectedPrompt {
                colortest.LogError(t,"Expected prompt message: "+expectedPrompt+ " got: "+ string(p))
            }else{
				colortest.LogSuccess(t, "TestPromptUsername completed successfully")
			}
            return len(p), nil
        },
    }

    // Create a new server instance
	srv := NewServer().(*Server)

    // Call the promptUsername function with the mock connection
    username := srv.promptUsername(mockConn)

    // Verify the username returned by the function
    expectedUsername := "layla"
    if username != expectedUsername {
        colortest.LogError(t,"Expected prompt message: "+expectedUsername+ " got: "+ username )
    }
}

 func TestSendReadyMessages(t *testing.T) {
    colortest.LogInfo(t, "Running TestSendReadyMessages...")
    srv := NewServer().(*Server)

    username := "layla"

	templateMessage := fmt.Sprintf("\n[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), username)

	mockConn := &mocks.MockConn{
		WriteFunc: func(p []byte) (n int, err error) {
            if string(p) != templateMessage{
				colortest.LogError(t,"Expcted ready message: "+templateMessage+" got: "+ string(p))
			} else{
				colortest.LogSuccess(t, "TestSendReadyMessages completed successfully")
			}

		return len(p), nil
		},
   }

   srv.sendReadyMessages(mockConn,username)
}

// TestMaximumClientLimitReached tests the scenario when the maximum client limit is reached.
func TestMaximumClientLimitReached(t *testing.T) {
	colortest.LogInfo(t, "Running TestMaximumClientLimitReached...")
    // Initialize server
	srv := NewServer().(*Server)
    addr := "localhost:9898"
    go func() {
        err := srv.InitializeServer(addr)
        if err != nil {
            t.Errorf( "\033[31m"+"Error initializing server: %v"+ "\033[0m", err)
			
        }else{
			colortest.LogSuccess(t, "TestMaximumClientLimitReached completed successfully")
		}
        err = srv.ListenAndServe()
        if err != nil {
			t.Errorf( "\033[31m"+"Error listening and serving: %v"+ "\033[0m", err)
        }else{
			colortest.LogSuccess(t, "TestMaximumClientLimitReached completed successfully")
		}
    }()
    
    // Simulate connections until the maximum client limit is reached
    maxClients := 10
    for i := 0; i < maxClients; i++ {
        conn, err := net.Dial("tcp", addr)
        if err != nil {
			t.Errorf( "\033[31m"+"Error connecting: %v"+ "\033[0m", err)
        }else{
			colortest.LogSuccess(t, "TestMaximumClientLimitReached completed successfully")
		}
        defer conn.Close()
        // Additional testing as needed
    }
    
    // Attempt to connect when the maximum client limit is reached
    _, err := net.Dial("tcp", addr)
    if err == nil {
		t.Errorf( "\033[31m"+"Expected error when maximum client limit reached, got %v"+ "\033[0m",err)
        return
    }
    expectedErrMsg := "Sorry, the chat room is full. Please try again later."
    if err.Error() != expectedErrMsg {
        t.Errorf( "\033[31m"+"Expected error message: %s, got: %s", expectedErrMsg, err.Error())
    }
}

// TestBroadcast tests the broadcast function.
// func TestBroadcast(t *testing.T) {
//     colortest.LogInfo(t, "Running TestBroadcast...")

//     // Create a new Server instance
//     srv := NewServer().(*Server)

// // Initialize the global Clients slice
// interfaces.Clients = []client.Client{
// 	{Name: "Client1", Conn: nil, Writer: bufio.NewWriter(nil)},
// 	{Name: "Client2", Conn: nil, Writer: bufio.NewWriter(nil)},
// }

// // Simulate a sender client
// sender := &client.Client{Name: "layla", Conn: nil, Writer: bufio.NewWriter(nil)}

// // Simulate broadcasting a message
// message := "Hello, world!"
// senderName := sender.Name
// messageType := "haha"
// time := time.Now().Format("2006-01-02 15:04:05")

// // Call the broadcast function
// // Assuming you have a function named broadcast in server.go that takes the necessary parameters
// srv.broadcast(message, senderName, sender.Conn, messageType)
// formattedMessage := fmt.Sprintf("\n[%s][%s]: %s\n", time, senderName, message)

// // Verify that the message was sent to all clients except the sender
// for _, client := range interfaces.Clients {
// 	if client.Name != senderName {
// 		// Here, we would typically check the contents of the client's writer.
// 		// However, since we're using a mock writer, we'll just print a message.
// 		fmt.Printf("Message sent to %s %s\n", client.Name,formattedMessage)
// 	}
// }
// }

// TestAddClient tests the addClient function.
func TestAddClient(t *testing.T) {
	// Setup mock objects for testing
	// Call addClient with a mock net.Conn object and username
	// Assert that the client is added to the server
}

// TestRemoveClient tests the removeClient function.
func TestRemoveClient(t *testing.T) {
	// Setup mock objects for testing
	// Add a mock client to the server
	// Call removeClient with the mock client's connection
	// Assert that the client is removed from the server
}

// TestCleanupHistoryFile tests the CleanupHistoryFile function.
func TestCleanupHistoryFile(t *testing.T) {
	// Create a dummy history file with some content
	// Call CleanupHistoryFile
	// Assert that the history file is cleared
}


// TestErrorHandling tests error handling throughout the code.
func TestErrorHandling(t *testing.T) {
	// Test various error scenarios and ensure they are handled gracefully
}

// TestEdgeCases tests edge cases such as empty messages, invalid usernames, etc.
func TestEdgeCases(t *testing.T) {
	// Test edge cases to ensure robustness of the code
}

// TestConcurrency tests concurrency aspects, like mutexes for thread-safe access.
func TestConcurrency(t *testing.T) {
	// Test concurrency aspects of the code
}

// TestIntegration tests interaction between different components of the server.
func TestIntegration(t *testing.T) {
	// Test interaction between different components of the server
}


