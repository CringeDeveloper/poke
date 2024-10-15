package main

var commandsMap map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*Paths) error
}

func init() {
	commandsMap = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "snow next 20 maps",
			callback:    commandMap,
		},
		"map_b": {
			name:        "map_b",
			description: "snow previous 20 maps",
			callback:    commandMapB,
		},
	}
}
