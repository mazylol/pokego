package pokego

import (
	"encoding/json"
	"fmt"
	pokemonCache "github.com/mazylol/pokego/cache/pokemon"
	"github.com/mazylol/pokego/types/pokemon"
	"log"
)

// GetPokemonAbility returns an Ability struct containing information about the Ability with the given name.
func GetPokemonAbility(name string) (pokemon.Ability, error) {
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

// GetPokemonCharacteristic returns a Characteristic struct containing information about the Characteristic with the given id.
func GetPokemonCharacteristic(id int) (pokemon.Characteristic, error) {
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

// GetPokemonEggGroup returns an EggGroup struct containing information about the EggGroup with the given name.
func GetPokemonEggGroup(name string) (pokemon.EggGroup, error) {
	eggGroup, err := pokemonCache.GetEggGroupFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("egg-group/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &eggGroup)

		if err == nil {
			err = pokemonCache.AddEggGroupToCache(eggGroup)
		}

		return eggGroup, err
	} else {
		return eggGroup, err
	}
}

// GetPokemonGender returns a Gender struct containing information about the EggGroup with the given name.
func GetPokemonGender(name string) (pokemon.Gender, error) {
	gender, err := pokemonCache.GetGenderFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("gender/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &gender)

		if err == nil {
			err = pokemonCache.AddGenderToCache(gender)
		}

		return gender, err
	} else {
		return gender, err
	}
}

// GetPokemonGrowthRate returns a GrowthRate struct containing information about the GrowthRate with the given name.
func GetPokemonGrowthRate(name string) (pokemon.GrowthRate, error) {
	growthRate, err := pokemonCache.GetGrowthRateFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("growth-rate/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &growthRate)

		if err == nil {
			err = pokemonCache.AddGrowthRateToCache(growthRate)
		}

		return growthRate, err
	} else {
		return growthRate, err
	}
}

// GetPokemonNature returns a Nature struct containing information about the Nature with the given name.
func GetPokemonNature(name string) (pokemon.Nature, error) {
	nature, err := pokemonCache.GetNatureFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("nature/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &nature)

		if err == nil {
			err = pokemonCache.AddNatureToCache(nature)
		}

		return nature, err
	} else {
		return nature, err
	}
}

// GetPokemonPokeathlonStats returns a PokeathlonStat struct containing information about the PokeathlonStat with the given name.
func GetPokemonPokeathlonStats(name string) (pokemon.PokeathlonStat, error) {
	pokeathlonStat, err := pokemonCache.GetPokeathlonStatFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokeathlon-stat/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &pokeathlonStat)

		if err == nil {
			err = pokemonCache.AddPokeathlonStatToCache(pokeathlonStat)
		}

		return pokeathlonStat, err
	} else {
		return pokeathlonStat, err
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

// GetPokemonLocationArea returns a PokemonLocationArea struct containing information about the PokemonLocationArea with the given name.
func GetPokemonLocationArea(name string) ([]pokemon.LocationArea, error) {
	locationArea, err := pokemonCache.GetLocationAreaFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon/%v/encounters", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &locationArea)

		if err == nil {
			err = pokemonCache.AddLocationAreaToCache(locationArea, name)
		}

		return locationArea, err
	} else {
		return locationArea, err
	}
}

// GetPokemonColor returns a Color struct containing information about the Color with the given name.
func GetPokemonColor(name string) (pokemon.Color, error) {
	color, err := pokemonCache.GetColorFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon-color/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &color)

		if err == nil {
			err = pokemonCache.AddColorToCache(color)
		}

		return color, err
	} else {
		return color, err
	}
}

// GetPokemonForm returns a Form struct containing information about the Form with the given name.
func GetPokemonForm(name string) (pokemon.Form, error) {
	form, err := pokemonCache.GetFormFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon-form/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &form)

		if err == nil {
			err = pokemonCache.AddFormToCache(form)
		}

		return form, err
	} else {
		return form, err
	}
}

// GetPokemonHabitat returns a Habitat struct containing information about the Habitat with the given name.
func GetPokemonHabitat(name string) (pokemon.Habitat, error) {
	habitat, err := pokemonCache.GetHabitatFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon-habitat/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &habitat)

		if err == nil {
			err = pokemonCache.AddHabitatToCache(habitat)
		}

		return habitat, err
	} else {
		return habitat, err
	}
}

// GetPokemonShape returns a Shape struct containing information about the Shape with the given name.
func GetPokemonShape(name string) (pokemon.Shape, error) {
	shape, err := pokemonCache.GetShapeFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon-shape/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &shape)

		if err == nil {
			err = pokemonCache.AddShapeToCache(shape)
		}

		return shape, err
	} else {
		return shape, err
	}
}

// GetPokemonSpecies returns a Species struct containing information about the Species with the given name.
func GetPokemonSpecies(name string) (pokemon.Species, error) {
	species, err := pokemonCache.GetSpeciesFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("pokemon-species/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &species)

		if err == nil {
			err = pokemonCache.AddSpeciesToCache(species)
		}

		return species, err
	} else {
		return species, err
	}
}

// GetPokemonStat returns a Stat struct containing information about the Stat with the given name.
func GetPokemonStat(name string) (pokemon.Stat, error) {
	stat, err := pokemonCache.GetStatFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("stat/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &stat)

		if err == nil {
			err = pokemonCache.AddStatToCache(stat)
		}

		return stat, err
	} else {
		return stat, err
	}
}

// GetPokemonType returns a Type struct containing information about the Type with the given name.
func GetPokemonType(name string) (pokemon.Type, error) {
	tipe, err := pokemonCache.GetTypeFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("type/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &tipe)

		if err == nil {
			err = pokemonCache.AddTypeToCache(tipe)
		}

		return tipe, err
	} else {
		return tipe, err
	}
}
