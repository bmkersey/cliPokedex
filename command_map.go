package main

import (
	"errors"
	"fmt"
)

func commandMapf(c *config) error {
	locationsResp, err := c.pokeapiClient.ListLocations(c.next)
	if err != nil {
		return err
	}
	c.next = locationsResp.Next
	c.previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(c *config) error {
	if c.previous == nil {
		return errors.New("No previous pages yet.")
	}

	locationResp, err := c.pokeapiClient.ListLocations(c.previous)
	if err != nil {
		return err
	}
	c.next = locationResp.Next
	c.previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
