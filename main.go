package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Print "Enter text: "
	fmt.Print("Enter text: ")

	// 2. Create a scanner that reads from os.Stdin.
	scanner := bufio.NewScanner(os.Stdin)

	// 3. Read one line.
	scanner.Scan()

	// 4. Normalize and split the input.
	text := scanner.Text()

	text = strings.ToLower(text)
	words := strings.Fields(text)
	fmt.Printf("%q\n", words)

	result := map[string]int{}

	for _, word := range words {
		result[word]++
	}

	fmt.Println(result)
}
