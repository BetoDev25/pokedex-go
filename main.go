package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	nextURL     *string
	previousURL *string
}

type cliCommand struct {
	name 	    string
	description string
	callback    func(*config) error
}

var commandList map[string]cliCommand

func main() {

	commandList = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: 	     "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}

	reader := bufio.NewScanner(os.Stdin)

	cfg := &config{}
	for {
		fmt.Print("Pokedex > ")

		if !reader.Scan() {
			return
		}

		line := CleanInput(reader.Text())
		if len(line) == 0 {
			continue
		}

		command := line[0]
		cmd, exists := commandList[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}
