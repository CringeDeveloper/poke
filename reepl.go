package main

import (
	"fmt"
	"log"
	"pokedex/internal/commands"
)

func startReepl() {
	paths := commands.NewPaths()

	var command string
	var arg string
	for {
		fmt.Print("Pokedex > ")
		fmt.Scan(&command, &arg)
		if val, ok := commands.CommandsMap[command]; ok {
			err := val.Callback(paths, arg)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
