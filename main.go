package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		if command, exists := commands[input]; exists {
			err := command.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Command not recognized")
		}
	}

}

func commandHelp() error {
	fmt.Println("Welcome to the CLI Pokedex")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("Thank you for using the Pokedex.")
	os.Exit(0)
	return nil
}
