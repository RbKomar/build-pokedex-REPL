package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]
	exploreResp, err := cfg.pokeapiClient.ExploreLocation(locationName)
	if err != nil {
		return err
	}
	for _, encounter := range exploreResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
