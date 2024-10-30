package main

import "fmt"

func commandMap(cfg *config, arg string) error {
	loc, err := cfg.pokeapiClient.GetMap(cfg.nextLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = loc.Next
	cfg.prevLocation = loc.Previous

	for _, loc := range loc.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config, arg string) error {
	loc, err := cfg.pokeapiClient.GetMap(cfg.prevLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = loc.Next
	cfg.prevLocation = loc.Previous

	for _, loc := range loc.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
