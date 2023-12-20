package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/samvimes01/go-bootdev-pokedexcli/internal/commands"
	"github.com/samvimes01/go-bootdev-pokedexcli/internal/utils"
)

func StartRepl() {
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
		command, ok := commands.GetCommandsMap()[cmd[0]]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		command.Callback()
	}
}
