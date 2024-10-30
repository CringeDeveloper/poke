package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func startReepl(cfg *config) {
	var arg string
	for {
		fmt.Print("Pokedex > ")

		in := bufio.NewReader(os.Stdin)
		command, _ := in.ReadString('\n')
		command = strings.TrimSpace(command)
		parsed := strings.Split(command, " ")
		if len(parsed) == 1 {
			command = parsed[0]
			arg = ""
		} else if len(parsed) > 1 {
			command = parsed[0]
			arg = parsed[1]
		} else {
			command = ""
			arg = ""
		}

		if val, ok := CommandsMap[command]; ok {
			err := val.Callback(cfg, arg)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
