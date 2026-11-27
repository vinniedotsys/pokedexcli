package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/vinniedotsys/pokedexcli/internal/pokeapi"
)


type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}


type config struct {
	commands map[string]cliCommand
	location_next string
	location_previous string
	pokedex map[string]pokeapi.Pokemon
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
		    "map": {
			name:	     "map",
			description: "List 20 locations areas in the Pokemon world",
			callback: commandMap,
			
		    },
		    "mapb": {
			name:	     "mapb",
			description: "List the previous 20 locations areas in the Pokemon world",
			callback: commandMapBack,
			
		    },
		   "explore": {
			name:	    "explore",
			description: "Display the names of all Pokémon found in the area passed in parameter",
			callback: commandExplore,
		   },
		   "catch": {
			name:	    "catch",
			description: "Tries to catch to pokemon passed in parameter, if successfull, adds it to the Pokedex",
			callback: commandCatch,
		   },
		   "inspect": {
			name:	    "inspect",
			description: "Gives information about a captured Pokemon",
			callback: commandInspect,
		   },
		   "pokedex": {
			name:	    "pokedex",
			description: "List Pokemons in your Pokedex",
			callback: commandPokedex,
		   },
			
		}

}


func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {

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
			args := input[1:]
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println("Error executing command:", err)
			}

		}
		
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading input: %v", err)
		}

	}



}
