package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name 	    string
	description string
	callback    func() error
}

var commandList map[string]cliCommand

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range commandList {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

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
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		var line []string

		if reader.Scan() {
			userInput := reader.Text()
			line = CleanInput(userInput)
		}
		command := line[0]
		if command != "" {
			if _, exists := commandList[command]; exists {
				commandList[command].callback()
			} else {
				fmt.Print("Unknown command")
			}
		}
	}
}
