package main

import (
	"github.com/vinniedotsys/pokedexcli/internal/pokeapi"
	"fmt"
)





func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Command need 1 argument when %v where provided", len(args))
	}

	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	result, err := pokeapi.GetSpecificArea(url)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + args[0] + "...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range result.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}

