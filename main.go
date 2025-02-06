package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func cleanInput(text string) []string {
	substrings := strings.Fields(text)
	for i, str := range substrings {
		substrings[i] = strings.TrimSpace(strings.ToLower(str))
	}
	return substrings

}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) > 0 {
			fmt.Printf("Your command was: %s\n", cleanedInput[0])
		}
	}
}
