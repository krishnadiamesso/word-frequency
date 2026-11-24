package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. Print "Enter text: "
	fmt.Print("Enter text: ")

	// 2. Create a scanner that reads from os.Stdin.
	scanner := bufio.NewScanner(os.Stdin)

	// 3. Read one line.
	scanner.Scan()

	// 4. Print the line you read.
	fmt.Println("You entered: ", scanner.Text())
}
