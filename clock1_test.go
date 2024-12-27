package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

func TestClockServer(t *testing.T) {
	// Start the clock server in a goroutine
	go func() {
		main()
	}()

	// Give the server some time to start
	time.Sleep(1 * time.Second)

	// Connect to the clock server
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Read data from the server
	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		t.Fatalf("Failed to read from server: %v", scanner.Err())
	}
	response := scanner.Text()

	// Verify the response format
	_, err = time.Parse("15:04:05", strings.TrimSpace(response))
	if err != nil {
		t.Errorf("Received invalid time format: %v", response)
	}
}
