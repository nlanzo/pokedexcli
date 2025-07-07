package main

import (
	"time"

	"github.com/nlanzo/pokedexcli/internal/pokeapi"
)

const (
	TIMEOUT = 5 * time.Second
	CACHE_INTERVAL = 10 * time.Second
)

func main() {
	pokeClient := pokeapi.NewClient(TIMEOUT, CACHE_INTERVAL)
	config := &Config{
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}