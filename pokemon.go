package pokego

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mazylol/pokego/cache"
	"github.com/mazylol/pokego/types"
)

// GetPokemon returns a Pokemon struct containing information about the Pokemon with the given name.
func GetPokemon(name string) (types.Pokemon, error) {
	pokemon, err := cache.GetPokemonFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &pokemon)

		if err == nil {
			err = cache.AddPokemonToCache(pokemon)
		}

		return pokemon, err
	} else {
		return pokemon, err
	}
}

// GetPokemonList returns a list of Pokemon names. You have to include a limit for the amount of names you want.
func GetPokemonList(limit int) (types.PokemonList, error) {
	pokemonList, err := cache.GetPokemonListFromCache(limit)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon?limit=%v", limit))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &pokemonList)

		if err == nil {
			err = cache.AddPokemonListToCache(pokemonList, limit)
		}

		return pokemonList, err
	} else {
		return pokemonList, err
	}
}
