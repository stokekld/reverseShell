package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	//var line string

	fmt.Println("Initializing server...")

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error initializing server")
		os.Exit(1)
	}

	conn, err := ln.Accept()

	if err != nil {
		// handle error
	}

	go func() { io.Copy(os.Stdout, conn) }()

	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	io.Copy(conn, os.Stdin)

	terminal.Restore(int(os.Stdin.Fd()), oldState)
}
