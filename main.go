package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"sync"
)

type Result struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type Pokemon struct {
	Forms []struct {
		Name string `json:"name"`
	} `json:"forms"`
	GameIndicies []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
		} `json:"version"`
	} `json:"game_indices"`
}

const API = "https://pokeapi.co/api/v2/"

func getPokemons() *Result {
	resp, err := http.Get(API + "pokemon/?limit=100")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	var pokemons Result
	err = json.NewDecoder(resp.Body).Decode(&pokemons)
	if err != nil {
		fmt.Println(err)
	}
	return &pokemons

}

func getPokemon(name string) *Pokemon {
	resp, err := http.Get(API + "pokemon/" + name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		fmt.Println(err)
	}
	return &pokemon
}

func main() {
	names := getPokemons()
	var wg sync.WaitGroup

	for _, name := range names.Results {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			p := getPokemon(n)
			if p != nil {
				fmt.Printf("Pokemon name: %s \nGame index: %d \nVersion: %s \n\n\n", p.Forms[0].Name, p.GameIndicies[0].GameIndex, p.GameIndicies[rand.IntN(len(p.GameIndicies))].Version.Name)
			}
		}(name.Name)
	}
	wg.Wait()
}
