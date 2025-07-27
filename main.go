package main

import (
	"github.com/seiobata/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient()
	cfg := &config{
		Client: client,
	}
	startRepl(cfg)
}
