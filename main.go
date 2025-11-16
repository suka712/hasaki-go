package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hello world!")

	runCmd("git", "add", ".")

	commitMsg := "Auto generated: Made changes to the code."
	runCmd("git", "commit", "-m", commitMsg)
}

func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run();
}