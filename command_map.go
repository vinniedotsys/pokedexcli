package main

import (
	"github.com/vinniedotsys/pokedexcli/internal/pokeapi"
	"fmt"
)





func commandMap(cfg *config) error {
	var url string
	if cfg.location_next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = cfg.location_next
	}

	result, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.location_previous = cfg.location_next
	if result.Next != nil {
		cfg.location_next = *result.Next
	} else {
		cfg.location_next = ""
	}

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}
	return nil
}


func commandMapBack(cfg *config) error {
	var url string
	if cfg.location_previous == "" {
		fmt.Println("You're on the first page")
		return nil
	} else {
		url = cfg.location_previous
	}

	result, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.location_next = cfg.location_previous
	if result.Previous != nil {
		cfg.location_previous = *result.Previous
	} else {
		cfg.location_previous = ""
	}

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}
	return nil
}

