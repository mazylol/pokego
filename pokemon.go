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
	poke, err := pokemon2.GetPokemonFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &poke)

		if err == nil {
			err = pokemon2.AddPokemonToCache(poke)
		}

		return poke, err
	} else {
		return poke, err
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

// GetAbility returns an Ability struct containing information about the Ability with the given name.
func GetAbility(name string) (pokemon.Ability, error) {
	ability, err := pokemon2.GetAbilityFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("ability/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &ability)

		if err == nil {
			err = pokemon2.AddAbilityToCache(ability)
		}

		return ability, err
	} else {
		return ability, err
	}
}

// GetAbilityList returns a list of Ability names. You have to include a limit for the amount of names you want.
func GetAbilityList(limit int) (pokemon.AbilityList, error) {
	abilityList, err := pokemon2.GetAbilityListFromCache(limit)
	if err != nil {
		body, err := callApi(fmt.Sprintf("ability?limit=%v", limit))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &abilityList)

		if err == nil {
			err = pokemon2.AddAbilityListToCache(abilityList, limit)
		}

		return abilityList, err
	} else {
		return abilityList, err
	}
}
