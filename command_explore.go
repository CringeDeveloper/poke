package main

import (
	"fmt"
)

func commandExplore(cfg *config, arg string) error {
	area, err := cfg.pokeapiClient.GetExplore(arg)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", arg)
	fmt.Println("Found Pokemon:")
	area.PrintPokemon()

	return nil
}
