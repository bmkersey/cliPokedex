package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bmkersey/cliPokedex/internal/pokeapi"
)

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		command := input[0]

		if command, exists := getCommands()[command]; exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println("Command not recognized")
			continue
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
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
			description: "Retrive a list of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Retrieve the previous page of locations",
			callback:    commandMapb,
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
