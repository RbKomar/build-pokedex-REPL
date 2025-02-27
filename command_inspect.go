package main

import (
	"errors"
	"fmt"
	"pokedexcli/internal/pokeapi"
	"strconv"
)

func printPokemon(pokemon pokeapi.Pokemon) {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, statElem := range pokemon.Stats {
		fmt.Printf("  -%s: %s\n", statElem.Stat.Name, strconv.Itoa(statElem.BaseStat))
	}
	fmt.Println("Types:")
	for _, typesElem := range pokemon.Types {
		fmt.Printf("  - %s\n", typesElem.Type.Name)
	}
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("missing pokemon name")
	}
	name := args[0]

	pokemon, exists := cfg.caughtPokemons[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	printPokemon(pokemon)
	return nil
}
