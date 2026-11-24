package main

import (
	"os"
	"fmt"
	)




func commandExit(cfg *config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
