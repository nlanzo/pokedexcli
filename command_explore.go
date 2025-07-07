package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a location name")
	}

	location := args[0]

	locationRes, err := config.pokeapiClient.GetLocation(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationRes.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationRes.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}