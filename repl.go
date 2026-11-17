package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := cleanInput(scanner.Text())
			fmt.Printf("Your command was: %v\n", input[0])
		}
		
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error readind input: %v", err)
		}

	}



}
