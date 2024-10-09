package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Println()
	fmt.Println("Welcome to the CLI Pokedex")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}
