package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"math/rand"
	"time"
)

/*
type Pokemon struct {
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
}
*/

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("pokemon name required")
	}

	name := args[0]

	var url string
	url = "https://pokeapi.co/api/v2/pokemon/" + name

	body, ok := cfg.cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Unknown error")
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return fmt.Errorf("Error: status code %d, %s pokemon not found", res.StatusCode, name)
		}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Unknown error")
		}

		cfg.cache.Add(url, body)
	}

	//unmarshal
	var poke Pokemon
	if err := json.Unmarshal(body, &poke); err != nil {
		return fmt.Errorf("Could not parse JSON")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", poke.Name)
	time.Sleep(1 * time.Second)

	roll := rand.Float64() * 380.0
	if roll > float64(poke.BaseExperience) {
		cfg.pokedex[poke.Name] = poke
		fmt.Printf("%s was caught!\n", poke.Name)
	} else {
		fmt.Printf("%s escaped!\n", poke.Name)
	}
	return nil
}
