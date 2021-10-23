package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)

	switch input {
	case "hun\r\n":
		fmt.Println("Welcome hun!")
	case "zi\r\n":
		fmt.Println("Welcome zi!")
	case "qin\r\n":
		fmt.Println("Welcome qin!")
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}

	switch input {
	case "hun\r\n":
		fallthrough
	case "zi\r\n":
		fallthrough
	case "qin\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	switch input {
	case "hun\r\n", "zi\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}
