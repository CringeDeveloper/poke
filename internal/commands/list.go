package commands

var CommandsMap map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	Callback    func(*Paths, string) error
}

func init() {
	CommandsMap = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "snow next 20 maps",
			Callback:    commandMap,
		},
		"map_b": {
			name:        "map_b",
			description: "snow previous 20 maps",
			Callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "explore the <area_name>",
			Callback:    commandExplore,
		},
	}
}
