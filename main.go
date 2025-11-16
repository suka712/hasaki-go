package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	runCmd("git", "add", ".")

	filesChanged := getChanges()
	if len(filesChanged) == 0 {
		fmt.Println("No file change. No commit made.")
		return
	}

	diff := getDiff()
	if len(diff) == 0 {
		fmt.Println("Diff output is empty. No commit made.")
		return
	}

	// commitMsg := "Auto generated: Made changes to the code."
	// runCmd("git", "commit", "-m", commitMsg)
}

// ---------------------------Shell & Git helpers---------------------------
func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func getChanges() []string {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting changed files:", err)
		return nil
	}

	if strings.TrimSpace(string(out)) == "" {
		return []string{}
	}

	return strings.Split(strings.TrimSpace(string(out)), "\n")
}

func getDiff() string {
	cmd := exec.Command("git", "--no-pager", "diff", "--staged")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting diff:", err)
		return ""
	}
    
	return string(out)
}
