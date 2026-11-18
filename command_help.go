package main

import (
	"fmt"
	)




func commandHelp() error {

	commands := commandsList()
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, detail := range commands {
		fmt.Printf("%v: %v\n", detail.name, detail.description)
	}
	return  nil

}
