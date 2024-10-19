package main

import (
	"fmt"
	"log"
	"pokedex/internal/commands"
)

func startReepl() {
	paths := commands.NewPaths()

	var userInput string
	for {
		fmt.Print("Pokedex > ")
		fmt.Scan(&userInput)
		if val, ok := commands.CommandsMap[userInput]; ok {
			err := val.Callback(paths)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
