package main

import (
	"fmt"
	"os"
)

func commandExit(_ *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}