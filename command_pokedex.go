package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	for _, p := range cfg.caughtPokemons {
		fmt.Println("  -", p.Name)
	}
	return nil
}
