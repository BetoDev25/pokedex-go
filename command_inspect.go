package main

import (
	"fmt"
)

type Pokemon struct {
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("pokemon name required")
	}

	name := args[0]

	poke, ok := cfg.pokedex[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Height: %d\n", poke.Height)
	fmt.Printf("Weight: %d\n", poke.Weight)

	fmt.Println("Stats:")
	for _, s := range poke.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
