package pokego

import (
	"encoding/json"
	"fmt"
	pokemonCache "github.com/mazylol/pokego/cache/pokemon"
	"github.com/mazylol/pokego/types/pokemon"
	"log"
)

// GetAbility returns an Ability struct containing information about the Ability with the given name.
func GetAbility(name string) (pokemon.Ability, error) {
	ability, err := pokemonCache.GetAbilityFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("ability/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &ability)

		if err == nil {
			err = pokemonCache.AddAbilityToCache(ability)
		}

		return ability, err
	} else {
		return ability, err
	}
}

// GetCharacteristic returns a Characteristic struct containing information about the Characteristic with the given id.
func GetCharacteristic(id int) (pokemon.Characteristic, error) {
	characteristic, err := pokemonCache.GetCharacteristicFromCache(id)
	if err != nil {
		body, err := callApi(fmt.Sprintf("characteristic/%v", id))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &characteristic)

		if err == nil {
			err = pokemonCache.AddCharacteristicToCache(characteristic)
		}

		return characteristic, err
	} else {
		return characteristic, err
	}
}

// GetPokemon returns a Pokemon struct containing information about the Pokemon with the given name.
func GetPokemon(name string) (pokemon.Pokemon, error) {
	poke, err := pokemonCache.GetPokemonFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &poke)

		if err == nil {
			err = pokemonCache.AddPokemonToCache(poke)
		}

		return poke, err
	} else {
		return poke, err
	}
}
