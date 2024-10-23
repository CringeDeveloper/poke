package commands

import (
	"bytes"
	"encoding/gob"
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

	var loc pokeapi.Locations
	var myBuffer bytes.Buffer

	enc := gob.NewEncoder(&myBuffer)

	if val, ok := cacheMap.Get(*position); ok {
		fmt.Println("WE ARE HERE!!!!!!")
		reader := bytes.NewReader(val)
		dec := gob.NewDecoder(reader)
		err := dec.Decode(&loc)

		if err != nil {
			log.Fatal("cannot decode")
		}
	} else {
		loc, _ = pokeapi.GetLocations(position)

		_ = enc.Encode(loc)

		cacheMap.Add(*position, myBuffer.Bytes())
	}

	paths.Next = loc.Next
	paths.Previous = loc.Previous
	pokeapi.PrintLocations(&loc)

	return nil
}
