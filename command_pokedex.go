package main

import (
	"fmt"
)

func commandPokedex(config *Config, args ...string) error {
	if len(config.caughtPokemon) == 0 { 
		fmt.Println("Your Pokedex is empty")
	} else {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range config.caughtPokemon {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	}
	return nil
}