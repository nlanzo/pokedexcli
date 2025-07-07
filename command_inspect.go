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
	fmt.Printf("  -hp: %v\n", pokemon.Stats[0].BaseStat)
	fmt.Printf("  -attack: %v\n", pokemon.Stats[1].BaseStat)
	fmt.Printf("  -defense: %v\n", pokemon.Stats[2].BaseStat)
	fmt.Printf("  -special-attack: %v\n", pokemon.Stats[3].BaseStat)
	fmt.Printf("  -special-defense: %v\n", pokemon.Stats[4].BaseStat)
	fmt.Printf("  -speed: %v\n", pokemon.Stats[5].BaseStat)

	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}