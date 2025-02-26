package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient:  &pokeClient,
		caughtPokemons: map[string]pokeapi.Pokemon{},
	}
	runRepl(cfg)
}
