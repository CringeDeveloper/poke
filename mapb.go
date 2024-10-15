package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func commandMapB(paths *Paths) error {
	if paths.Previous == nil {
		return fmt.Errorf("no previous locations")
	}

	res, err := http.Get(*paths.Previous)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var loc Locations

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&loc)
	if err != nil {
		log.Fatal(err)
	}

	paths.Next = loc.Next
	paths.Previous = loc.Previous
	printLocations(&loc)

	return nil
}
