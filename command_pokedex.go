package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	body := cfg.pokedex
	if len(body) == 0 {
		return fmt.Errorf("pokedex is empty!")
	}

	fmt.Println("Your Pokedex:")
	for poke := range body {
		fmt.Printf("  - %s\n", poke)
	}
	return nil
}
