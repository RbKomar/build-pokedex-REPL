package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("missing pokemon name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	randomNumber := rand.Intn(pokemon.BaseExperience)
	if randomNumber == 0 {
		fmt.Println(name, "was caught!")
		cfg.caughtPokemons[name] = pokemon
	} else {
		fmt.Println(name, "escaped!")
	}
	return nil
}
