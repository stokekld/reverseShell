package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("Initializing server...")

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error initializing server")
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			// handle error
		}

		fmt.Printf("%T\n", conn)

		conn.Write([]byte("Hello from server"))
	}
}
