package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type LocationArea struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
	err := GetLocations(cfg, false)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}
	err := GetLocations(cfg, true)
	if err != nil {
		return err
	}
	return nil
}

func GetLocations(cfg *config, back bool) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if back {
		if cfg.Previous != nil {
			url = *cfg.Previous
		}
	} else {
		if cfg.Next != nil {
			url = *cfg.Next
		}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	res, err := cfg.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	locations := LocationArea{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return err
	}
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
