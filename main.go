package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	for i := range 10 {
		if i % 2 == 0 {
			fmt.Println("The index we are at is divisible by 2")
		} else {
			fmt.Println("The index we are at not divisible by 2")
		}
	}
}
