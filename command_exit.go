package main

import (
	"fmt"
	"os"
)

func commandExit(config *config) error {
	fmt.Println("Thank you for using the Pokedex.")
	os.Exit(0)
	return nil
}
