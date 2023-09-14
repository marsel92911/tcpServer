package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	cfg := getConfig()
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", cfg.port))
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is running and listening on port 12345")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn, cfg)
	}
}
