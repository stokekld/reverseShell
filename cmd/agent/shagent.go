package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	buffer := make([]byte, 1024)

	fmt.Println("Initializing client...")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("Error initializing client")
		os.Exit(1)
	}

	fmt.Println("Connected")

	fmt.Printf("%T\n", conn)

	conn.Read(buffer)

	fmt.Println("Executing...", string(buffer))
}
