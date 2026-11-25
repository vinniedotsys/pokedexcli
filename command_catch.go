package main

import (
	"fmt"
	"github.com/vinniedotsys/pokedexcli/internal/pokeapi"
	"math/rand"
	"math"
)

func NormalizeExp(exp float64) float64 {
	return ((math.Log(exp) - math.Log(36)) / (math.Log(635) - math.Log(36))) * 100 // max XP 635 ; min XP 36
}


func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Command need 1 argument when %v where provided", len(args))
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])

	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]
	result, err := pokeapi.GetPokemon(url)
	if err != nil {
		return err
	}
	catchability := NormalizeExp(float64(result.BaseExperience))
	pokeball := rand.Float64() * 100

	if pokeball > catchability {
		fmt.Printf("%v was caught!\n", args[0])
		cfg.pokedex[args[0]] = result
	} else {
		fmt.Printf("%v escaped!\n", args[0])
	}
	return nil
}
