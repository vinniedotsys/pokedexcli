package main

import (
	"fmt"
	)




func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) < 1 {
		return fmt.Errorf("Your pokedex contains %v Pokemon", len(cfg.pokedex))
	}
	fmt.Println("Your Pokedex :")
	for pokemon, _ := range cfg.pokedex {
		fmt.Printf("  - %v\n", pokemon)
	}

	return nil
}
