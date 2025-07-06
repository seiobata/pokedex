package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Hello, World!\n")
}

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	new_text := strings.Fields(lower_text)
	return new_text
}
