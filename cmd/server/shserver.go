package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	
	var line string

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

	fmt.Printf("%T\n", conn)

	go func(){
		for {
			buffer := make([]byte, 1024)
			
			conn.Read(buffer)

			fmt.Print(string(buffer))
		}
	}()

	for {
		fmt.Scanln(&line)
		conn.Write([]byte(line + "\n"))	
		line = ""
	}
	

}