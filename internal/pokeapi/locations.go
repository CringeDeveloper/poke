package pokeapi

import "fmt"

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func PrintLocations(loc *Locations) {
	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}
}
