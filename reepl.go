package main

import (
	"fmt"
	"log"
	"pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func startReepl(cfg *config) {
	var command string
	var arg string
	for {
		fmt.Print("Pokedex > ")
		fmt.Scan(&command, &arg)
		if val, ok := CommandsMap[command]; ok {
			err := val.Callback(cfg, arg)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
