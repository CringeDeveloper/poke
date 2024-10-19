package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
)

var cacheMap = pokecache.NewCache() // TODO: think about it

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

	locBytes := new(bytes.Buffer) // TODO: change
	json.NewEncoder(locBytes).Encode(loc)
	cacheMap.Add(*position, locBytes.Bytes())

	paths.Next = loc.Next
	paths.Previous = loc.Previous
	pokeapi.PrintLocations(&loc)

	return nil
}
