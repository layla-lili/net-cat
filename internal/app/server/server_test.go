package server

import (
	"bytes"
	"log"
	colortest "netcat/internal/app/colorTest"
	mocks "netcat/internal/app/mocks"
	"netcat/internal/interfaces"
	"testing"
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
		t.Errorf("InitializeServer failed %v", err)
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
            colortest.LogSuccess(t,"TestSendWelcomeMessage complete sucessfully")
        }
        return len(p), nil
    },
}
    // Create a buffer to capture the output
    var buf bytes.Buffer
    srv.logger = log.New(&buf, "", 0)

    // Call sendWelcomeMessage with the mock connection
    srv.sendWelcomeMessage(mockConn)
}

// TestPromptUsername tests the promptUsername function.
// func TestPromptUsername(t *testing.T) {
//     colortest.LogInfo(t, "Running TestPromptUsername...")

//     srv:= NewServer().(*Server)
// 	//Setup mock objects for testing
//         mockConn := &mocks.MockConn{
//         WriteFunc: func(p []byte) (n int, err error) {
//         UN := interfaces.NamePrompt
//          if string(p) != UN {
//             colortest.LogError(t,"Expected Ready Message: "+UN+" got: "+string(p))
//         } else {
//             colortest.LogSuccess(t,"TestSendReadyMessages complete sucessfully")
//         }
//         return len(p), nil
//         },
//     }
// 	// Call promptUsername with a mock net.Conn object
//     srv.promptUsername(mockConn)
// 	// Assert that a valid username is returned
// }

// func TestSendReadyMessages(t *testing.T) {
//     colortest.LogInfo(t, "Running TestSendReadyMessages...")
//     srv := NewServer().(*Server)
   
//     mockConn := &mocks.MockConn{
//         WriteFunc: func(p []byte) (n int, err error) {
        
//          templateMessage := fmt.Sprintf("\n[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), UN)
//          if string(p)+"heheh" != templateMessage {
//             colortest.LogError(t,"Expected Ready Message: "+templateMessage+" got: "+string(p))
//         } else {
//             colortest.LogSuccess(t,"TestSendReadyMessages complete sucessfully")
//         }
//         return len(p), nil
//         },
//     }
//     UN := srv.promptUsername(mockConn)
// // Call sendWelcomeMessage with the mock connection
//     srv.sendReadyMessages(mockConn,UN)
// }

// TestHandleClientMessages tests the handleClientMessages function.
func TestHandleClientMessages(t *testing.T) {
	// Setup mock objects for testing
	// Call handleClientMessages with a mock net.Conn object
	// Assert that client messages are handled correctly
}

// TestBroadcast tests the broadcast function.
func TestBroadcast(t *testing.T) {
	// Setup mock objects for testing
	// Call broadcast with mock message and sender
	// Assert that the message is sent to other clients except the sender
}

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

// TestHandleConnection tests the handleConnection function.
// func TestHandleConnection(t *testing.T) {
//     colortest.LogInfo(t, "TestHandleConnection...")

//     // Create a mock net.Conn object
//     mockConn := &mocks.MockConn{
//         RemoteAddrFunc: func() net.Addr {
//             return &net.IPAddr{IP: []byte{127, 0, 0, 1}}
//         },
//         CloseFunc: func() error {
//             // Define behavior for closing the mock connection
//             return nil // Or any other desired behavior
//         },
//         Reader: bytes.NewBuffer(nil), // Initialize Reader field
//         Writer: bytes.NewBuffer(nil), // Initialize Writer field
//     }

//     // Mock the username input
//     username := "testUser"

//     // Send the mocked username input to the server
//     _, err := fmt.Fprintf(mockConn, "%s\n", username)
//     if err != nil {
//         t.Fatalf("Failed to send username input: %v", err)
//     }

//     srv := NewServer().(*Server) // Cast to Server type

//     // Use a channel to signal the completion of handleConnection
//     done := make(chan struct{})
//     go func() {
//         srv.handleConnection(mockConn)
//         close(done)
//     }()

//     // Wait for handleConnection to complete or timeout after a certain duration
//     select {
//     case <-done:
//         // handleConnection completed
//         colortest.LogSuccess(t, "handleConnection completed successfully")
//     case <-time.After(5 * time.Second):
//         t.Error("handleConnection did not complete within the expected time")
//     }

//     // Close the mock connection after the test completes
//     if err := mockConn.Close(); err != nil {
//         t.Fatalf("Failed to close mock connection: %v", err)
//     }
// }
