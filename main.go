package main

import (
	"time"

	"github.com/bmkersey/cliPokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}
