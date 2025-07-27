package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config) error {
	locations, err := cfg.Client.GetLocations(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}
	locations, err := cfg.Client.GetLocations(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
