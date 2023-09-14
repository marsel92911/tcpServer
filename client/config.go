package main

import (
	"fmt"
	"os"
)

const (
	defaultPort = "12345"
)

var (
	port string
)

func getPort() string {
	port = os.Getenv("TCP_SERVER_PORT")

	if port == "" {
		fmt.Printf("Server port is not set: using default port %s\n", defaultPort)
		port = defaultPort
	}
	return port
}
