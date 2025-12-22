package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io"

	"github.com/BetoDev25/pokedex-go/internal/pokeapi"
)

func commandMap(cfg *config, args []string) error {
        var url string
        if cfg.nextURL == nil {
                url = "https://pokeapi.co/api/v2/location-area?limit=20"
        } else {
                url = *cfg.nextURL
        }

	body, ok := cfg.cache.Get(url)

	if !ok {
        	res, err := http.Get(url)
        	if err != nil {
                	log.Fatal(err)
        	}
        	defer res.Body.Close()

        	body, err = io.ReadAll(res.Body)
        	if res.StatusCode > 299 {
                	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
        	}
        	if err != nil {
                	log.Fatal(err)
        	}

		cfg.cache.Add(url, body)
	}

        //unmarshal 'body'
	var resp pokeapi.LocationAreaListResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Fatal(err)
	}

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	cfg.nextURL = resp.Next
	cfg.previousURL = resp.Previous

	return nil
}

func commandMapb(cfg *config, args []string) error {
	var url string
	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = *cfg.previousURL
	}

	body, ok := cfg.cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		cfg.cache.Add(url, body)
	}

	//unmarshal 'body'
	var resp pokeapi.LocationAreaListResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Fatal(err)
	}

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	cfg.nextURL = resp.Next
	cfg.previousURL = resp.Previous

	return nil
}
