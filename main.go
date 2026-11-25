package main

import "github.com/vinniedotsys/pokedexcli/internal/pokeapi"

func main() {
	cfg := &config{
		commands: commandsList(),
		pokedex: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
