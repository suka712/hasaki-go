package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Hello world!")

	command := exec.Command("git", "--no-pager" ,"diff")

	output, err := command.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
