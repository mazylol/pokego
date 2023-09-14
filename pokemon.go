package pokego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemon returns a Pokemon struct containing information about the Pokemon with the given name.
func GetPokemon(name string) (Pokemon, error) {
	pokemon, err := getPokemonFromCache(name)
	if err != nil {
		var url = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", name)

		req, err := http.NewRequest("GET", url, nil)

		res, err := http.DefaultClient.Do(req)

		defer func(Body io.ReadCloser) {
			err = Body.Close()
		}(res.Body)

		body, err := io.ReadAll(res.Body)

		err = json.Unmarshal(body, &pokemon)

		if err == nil {
			err = addPokemonToCache(pokemon)
		}

		return pokemon, err
	} else {
		return pokemon, err
	}
}

// GetPokemonList returns a list of Pokemon names. You have to include a limit for the amount of names you want.
func GetPokemonList(limit int) (PokemonList, error) {
	pokemonList, err := getPokemonListFromCache(limit)
	if err != nil {
		var url = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?limit=%v", limit)

		req, err := http.NewRequest("GET", url, nil)

		res, err := http.DefaultClient.Do(req)

		defer func(Body io.ReadCloser) {
			err = Body.Close()
		}(res.Body)

		body, err := io.ReadAll(res.Body)

		err = json.Unmarshal(body, &pokemonList)

		if err == nil {
			err = addPokemonListToCache(pokemonList, limit)
		}

		return pokemonList, err
	} else {
		return pokemonList, err
	}
}
