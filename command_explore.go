package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("area name required")
	}
	area := args[0]

	var url string
	url = "https://pokeapi.co/api/v2/location-area/" + area

	body, ok := cfg.cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Unkown error")
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return fmt.Errorf("Error: status code %d, %s area not found", res.StatusCode, area)
		}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Unkown error")
		}

		cfg.cache.Add(url, body)
	}

	//unmarshal
	var loc LocationArea
	if err := json.Unmarshal(body, &loc); err != nil {
		return fmt.Errorf("Could not parse JSON")
	}

	for _, encounter := range loc.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
