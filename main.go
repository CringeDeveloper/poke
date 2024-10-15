package main

import (
	"fmt"
	"log"
)

func main() {
	paths := newPaths()
	var userInput string

	for {
		fmt.Print("Pokedex > ")
		fmt.Scan(&userInput)
		if val, ok := commandsMap[userInput]; ok {
			err := val.callback(paths)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
