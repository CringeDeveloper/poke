package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pokedex/internal/pokeapi"
)

func commandExplore(paths *Paths) error {
	res, err := http.Get(pokeapi.BaseUrl + "/location-area/canalave-city-area")

	if err != nil {
		log.Fatal("error")
	}

	defer res.Body.Close()
	var locationArea pokeapi.LocationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationArea)
	if err != nil {
		log.Fatal("error decode")
	}

	fmt.Println("Exploring {location_name}...")
	fmt.Println("Found Pokemon:")
	locationArea.PrintPokemon()

	return nil
}
