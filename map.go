package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func printLocations(loc *Locations) {
	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}
}

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

	res, err := http.Get(*position)

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
