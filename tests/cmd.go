package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("ls", "-la")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("Error executing command")
		os.Exit(1)
	}

	path, _ := exec.LookPath("ls")
	fmt.Println("Searching path...", path)

	cmd2 := exec.Cmd{
		Path: "/usr/bin/ls",
		Args: []string{
			"ls",
			"-la",
		},
	}
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr

	err = cmd2.Run()

	if err != nil {
		fmt.Println("Error executing command")
		os.Exit(1)
	}
}
