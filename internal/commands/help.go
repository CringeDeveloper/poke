package commands

import (
	"fmt"
)

func commandHelp(paths *Paths, arg string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, v := range CommandsMap {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}
