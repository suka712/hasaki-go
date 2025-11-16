package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Hello world!")

	runCommand("git status")
}

func runCommand(line string) {
	parts := strings.Fields(line)
	cmd := exec.Command(parts[0], parts[1:]...)
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.Run();
}