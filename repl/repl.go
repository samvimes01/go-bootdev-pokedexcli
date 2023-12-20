package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/samvimes01/go-bootdev-pokedexcli/internal/pokeapi"
	"github.com/samvimes01/go-bootdev-pokedexcli/utils"
)

type replConfig struct {
	PokeapiClient  pokeapi.Client
	NextOffset     int
	PreviousOffset int
	Limit          int
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *replConfig) error
}

func StartRepl() {
	cfg := replConfig{pokeapi.NewClient(), 0, 0, 20}
	scanner := bufio.NewScanner(os.Stdin)

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
	for {
		fmt.Print("Pokedex > ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		prompt := scanner.Text()

		cmd, err := utils.GetCmdFromPrompt(prompt)
		if err != nil {
			continue
		}
		command, ok := getCommandsMap()[cmd[0]]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		err = command.callback(&cfg)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
	}
}

func getCommandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Each subsequent call displays the names of next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
	}
}
