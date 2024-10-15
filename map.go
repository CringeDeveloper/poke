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
	if paths.Next == nil {
		return fmt.Errorf("no next locations")
	}

	res, err := http.Get(*paths.Next)
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
