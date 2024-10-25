package commands

import (
	"bytes"
	"encoding/gob"
	"log"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
)

var cacheMap = pokecache.NewCache()

func commandMap(paths *Paths, arg string) error {
	err := mapController(paths, paths.Next)

	return err
}

func commandMapB(paths *Paths, arg string) error {
	err := mapController(paths, paths.Previous)

	return err
}

func mapController(paths *Paths, pageURL *string) error {
	url := pokeapi.BaseUrl + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	var loc pokeapi.Locations

	if val, ok := cacheMap.Get(url); ok {
		dec := gob.NewDecoder(bytes.NewReader(val))
		err := dec.Decode(&loc)

		if err != nil {
			log.Fatal("cannot decode")
		}
	} else {
		var myBuffer bytes.Buffer
		enc := gob.NewEncoder(&myBuffer)
		loc, _ = pokeapi.GetLocations(url)

		_ = enc.Encode(loc)

		cacheMap.Add(url, myBuffer.Bytes())
	}

	paths.Next = loc.Next
	paths.Previous = loc.Previous
	pokeapi.PrintLocations(&loc)

	return nil
}
