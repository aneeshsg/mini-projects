package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//func main() {
	fmt.Println("Hello world!")
	var name string
	var age int

	in := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the quiz!")

	// Input name using fmt.Scan
	fmt.Println("Enter your name:")
	fmt.Scan(&name)

	// Input age using fmt.Scan
	fmt.Println("Enter your age:")
	_, err := fmt.Scan(&age)
	if err != nil {
		fmt.Println("Invalid age. Please enter a valid number.")
		return
	}
	in.ReadString('\n')
	if age < 13 {
		fmt.Println("You cannot play")
	} else {
		// Quiz question using in.ReadString
		fmt.Println("Which one is better - RTX 3080 or RTX 3090?")
		line, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		line = strings.TrimSpace(line) // Trim newline and whitespace

		if line == "RTX 3090" {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect")
		}
	}
}
