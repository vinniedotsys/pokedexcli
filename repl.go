package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandsList() map[string]cliCommand {
	return map[string]cliCommand{
		    "exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		    },
		    "help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		    },
			
		}

}


func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := cleanInput(scanner.Text())
			commands := commandsList()
			command, ok := commands[input[0]]; 
			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			err := command.callback()
			if err != nil {
				fmt.Println("Error executing command:", err)
			}

		}
		
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading input: %v", err)
		}

	}



}
