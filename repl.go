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
		inputs := cleanInput(scanner.Text())
		command := inputs[0]

		if cmd, ok := getCommands()[command]; ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
	err := scanner.Err()
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	new_text := strings.Fields(lower_text)
	return new_text
}
