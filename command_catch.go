package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}


	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	res := rand.Intn(pokemon.BaseExperience)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	config.caughtPokemon[pokemon.Name] = pokemon


	return nil
}