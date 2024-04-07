package mocks

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"time"
)

// MockConn is a mock implementation of net.Conn for testing purposes.
type MockConn struct {
	io.Reader
    io.Writer
    LocalAddrFunc  func() net.Addr
    RemoteAddrFunc func() net.Addr
    CloseFunc      func() error
    WriteFunc      func(p []byte) (n int, err error) // Define WriteFunc field
    ReadFunc      func(p []byte) (n int, err error) // Define ReadFunc field
	// New field for reading messages with template
	ReadTemplateFunc func(p []byte) (n int, err error)

}

// Close implements the net.Conn Close method.
func (c *MockConn) Close() error {
	if c == nil || c.CloseFunc == nil {
		log.Println("Error: Mock connection is nil or CloseFunc is nil")
		return nil
	}

	if c.CloseFunc != nil {
		return c.CloseFunc()
	}
	return nil
}

// LocalAddr implements the net.Conn LocalAddr method.
func (c *MockConn) LocalAddr() net.Addr {
	if c == nil || c.LocalAddrFunc == nil {
		log.Println("Error: Mock connection is nil or LocalAddrFunc is nil")
		return nil
	}

	if c.LocalAddrFunc != nil {
		return c.LocalAddrFunc()
	}
	return nil
}

// RemoteAddr implements the net.Conn RemoteAddr method.
func (c *MockConn) RemoteAddr() net.Addr {
	if c == nil || c.RemoteAddrFunc == nil {
		log.Println("Error: Mock connection is nil or RemoteAddrFunc is nil")
		return nil
	}

	if c.RemoteAddrFunc != nil {
		return c.RemoteAddrFunc()
	}
	return nil
}

// SetDeadline implements the net.Conn SetDeadline method.
func (c *MockConn) SetDeadline(t time.Time) error {
	return nil
}

// SetReadDeadline implements the net.Conn SetReadDeadline method.
func (c *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

// SetWriteDeadline implements the net.Conn SetWriteDeadline method.
func (c *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
// Write implements the net.Conn Write method using WriteFunc.
func (c *MockConn) Write(p []byte) (n int, err error) {
	if c == nil || c.WriteFunc == nil {
		log.Println("Error: Mock connection is nil or WriteFunc is nil")
		return 0, nil
	}

	return c.WriteFunc(p)
}

// Read implements the net.Conn Read method.
func (c *MockConn) Read(p []byte) (n int, err error) {
    if c == nil || c.Reader == nil {
        return 0, errors.New("MockConn: Reader is nil")
    }
    
    // Return a predefined input data
    input := []byte("layla\n")
    copy(p, input)
    return len(input), nil
}

// CustomRead implements a custom Read method for reading messages with template.
func (c *MockConn) CustomRead(p []byte, reader io.Reader) (n int, err error) {
    if reader == nil {
        return 0, errors.New("MockConn: Reader is nil")
    }

    // Return a predefined input data
    input := []byte("Test message\n")
    senderName := "Sender"
    expectedMessage := fmt.Sprintf("\n[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), senderName, input)

    copy(p, expectedMessage)

    return reader.Read(p)
}




