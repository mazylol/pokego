package pokego

import (
	"encoding/json"
	"fmt"
	pokemon2 "github.com/mazylol/pokego/cache/pokemon"
	"github.com/mazylol/pokego/types/pokemon"
	"log"
)

// GetPokemon returns a Pokemon struct containing information about the Pokemon with the given name.
func GetPokemon(name string) (pokemon.Pokemon, error) {
	pokemon, err := pokemon2.GetPokemonFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &pokemon)

		if err == nil {
			err = pokemon2.AddPokemonToCache(pokemon)
		}

		return pokemon, err
	} else {
		return pokemon, err
	}
}

// GetPokemonList returns a list of Pokemon names. You have to include a limit for the amount of names you want.
func GetPokemonList(limit int) (pokemon.PokemonList, error) {
	pokemonList, err := pokemon2.GetPokemonListFromCache(limit)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon?limit=%v", limit))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &pokemonList)

		if err == nil {
			err = pokemon2.AddPokemonListToCache(pokemonList, limit)
		}

		return pokemonList, err
	} else {
		return pokemonList, err
	}
}
