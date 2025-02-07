package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands()
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
