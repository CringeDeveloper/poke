package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pokedex/internal/pokeapi"
)

func commandExplore(paths *Paths, arg string) error {
	res, err := http.Get(pokeapi.BaseUrl + "/location-area/" + arg)

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

	fmt.Printf("Exploring %s...\n", arg)
	fmt.Println("Found Pokemon:")
	locationArea.PrintPokemon()

	return nil
}
