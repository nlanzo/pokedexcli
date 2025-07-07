package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, exists := config.caughtPokemon[pokemonName]
	if !exists {
		return fmt.Errorf("you have not caught %s", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}