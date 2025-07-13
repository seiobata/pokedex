package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := cleanInput(scanner.Text())
		if len(input) > 0 {
			fmt.Println("Your command was:", input[0])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	new_text := strings.Fields(lower_text)
	return new_text
}
