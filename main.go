package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"math/rand"

	"github.com/BetoDev25/pokedex-go/internal/pokecache"
)

func NewCache(interval time.Duration) *pokecache.Cache {
	return pokecache.NewCache(interval)
}


type config struct {
	nextURL     *string
	previousURL *string
	cache	    *pokecache.Cache
	pokedex     map[string]Pokemon
}

type cliCommand struct {
	name 	    string
	description string
	callback    func(*config, []string) error
}

var commandList map[string]cliCommand

func main() {
	rand.Seed(time.Now().UnixNano())

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
		"explore": {
			name:	     "explore",
			description: "Get the pokemon encounters in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try your luck at catching a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "display caught Pokemon stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "display contents of pokedex",
			callback:    commandPokedex,
		},
	}

	reader := bufio.NewScanner(os.Stdin)

	cfg := &config{
		cache:   pokecache.NewCache(5 * time.Second),
		pokedex: make(map[string]Pokemon),
	}
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
		args := []string{}
		if len(line) > 1 {
			args = line[1:]
		}
		cmd, exists := commandList[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg, args); err != nil {
			fmt.Println(err)
		}
	}
}
