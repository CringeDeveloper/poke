package main

import (
	"fmt"
	"log"
	"pokedex/internal/pokeapi"
)

func commandMap(paths *Paths) error {
	err := mapController(paths, paths.Next)

	return err
}

func commandMapB(paths *Paths) error {
	err := mapController(paths, paths.Previous)

	return err
}

func mapController(paths *Paths, position *string) error {
	if position == nil {
		return fmt.Errorf("no such locations")
	}

	loc, err := pokeapi.GetLocations(position)

	if err != nil {
		log.Fatal(err)
	}

	paths.Next = loc.Next
	paths.Previous = loc.Previous
	pokeapi.PrintLocations(&loc)

	return nil
}
