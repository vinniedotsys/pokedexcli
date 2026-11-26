package main

import (
	"fmt"
	)




func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Command need 1 argument when %v where provided", len(args))
	}
	
	pokemon, ok := cfg.pokedex[args[0]]
	if !ok {
		return fmt.Errorf("%v was not caught yet !\n", args[0])
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, poketype := range pokemon.Types {
		fmt.Printf("  - %v\n", poketype.Type.Name)
	}
	return  nil

}
