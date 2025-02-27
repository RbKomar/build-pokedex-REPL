package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type config struct {
	pokeapiClient    *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemons   map[string]pokeapi.Pokemon
}

func runRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Println("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) > 0 {
			commandName := words[0]
			args := []string{}
			if len(words) > 1 {
				args = words[1:]
			}

			command, exists := commands[commandName]
			if exists {
				if err := command.callback(cfg, args...); err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Get the list of pokemons in given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect the pokemon",
			callback:    commandInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
