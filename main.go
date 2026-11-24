package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// 1. Print "Enter text: "
	fmt.Print("Enter text: ")

	// 2. Create a scanner that reads from os.Stdin.
	scanner := bufio.NewScanner(os.Stdin)

	// 3. Read one line.
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}
		return
	}

	// 4. Normalize and split the input.
	text := scanner.Text()

	text = strings.ToLower(text)
	words := strings.Fields(text)

	result := make(map[string]int)

	for _, word := range words {
		result[word]++
	}

	keys := make([]string, 0)

	for key := range result {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, word := range keys {
		fmt.Printf("%s: %d\n", word, result[word])
	}
}
