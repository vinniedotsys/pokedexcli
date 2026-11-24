package main

import (
	"fmt"
	)




func commandHelp(cfg *config, args ...string) error {

	commands := cfg.commands
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, detail := range commands {
		fmt.Printf("%v: %v\n", detail.name, detail.description)
	}
	return  nil

}
