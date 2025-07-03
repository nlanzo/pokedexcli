package main

import (
	"errors"
	"fmt"
)

func commandMap(config *Config) error {
	locationsRes, err := config.pokeapiClient.ListLocations(&config.mapNextURL)
	if err != nil {
		return err
	}

	if locationsRes.Next != nil {
		config.mapNextURL = *locationsRes.Next
	} else {
		config.mapNextURL = ""
	}
	
	if locationsRes.Previous != nil {
		config.mapPreviousURL = *locationsRes.Previous
	} else {
		config.mapPreviousURL = ""
	}

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.mapPreviousURL == "" {
		return errors.New("you're on the first page")
	}

	locationRes, err := config.pokeapiClient.ListLocations(&config.mapPreviousURL)
	if err != nil {
		return err
	}

	if locationRes.Next != nil {
		config.mapNextURL = *locationRes.Next
	} else {
		config.mapNextURL = ""
	}
	
	if locationRes.Previous != nil {
		config.mapPreviousURL = *locationRes.Previous
	} else {
		config.mapPreviousURL = ""
	}

	for _, loc := range locationRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}