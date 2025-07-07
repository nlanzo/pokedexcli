package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nlanzo/pokedexcli/internal/pokeapi"
)

type Config struct {
	caughtPokemon map[string]pokeapi.Pokemon
	pokeapiClient pokeapi.Client
	mapPreviousURL string
	mapNextURL string
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		commandName := cleanedInput[0]
		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command: ", commandName)
			continue
		}
	}
}


func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}



type cliCommand struct {
	name string
	description string
	callback func(config *Config, args ...string) error
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
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous 20 map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "List pokemon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all your caught pokemon",
			callback:    commandPokedex,
		},
	}
}