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

// GetEggGroup returns an EggGroup struct containing information about the EggGroup with the given name.
func GetEggGroup(name string) (pokemon.EggGroup, error) {
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

// GetGender returns a Gender struct containing information about the EggGroup with the given name.
func GetGender(name string) (pokemon.Gender, error) {
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

// GetGrowthRate returns a GrowthRate struct containing information about the GrowthRate with the given name.
func GetGrowthRate(name string) (pokemon.GrowthRate, error) {
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

// GetNature returns a Nature struct containing information about the Nature with the given name.
func GetNature(name string) (pokemon.Nature, error) {
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

// GetPokeathlonStat returns a PokeathlonStat struct containing information about the PokeathlonStat with the given name.
func GetPokeathlonStat(name string) (pokemon.PokeathlonStat, error) {
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

// GetLocationArea GetPokemonLocationArea returns a PokemonLocationArea struct containing information about the PokemonLocationArea with the given name.
func GetLocationArea(name string) ([]pokemon.LocationArea, error) {
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

// GetColor returns a Color struct containing information about the Color with the given name.
func GetColor(name string) (pokemon.Color, error) {
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

// GetForm returns a Form struct containing information about the Form with the given name.
func GetForm(name string) (pokemon.Form, error) {
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

// GetHabitat returns a Habitat struct containing information about the Habitat with the given name.
func GetHabitat(name string) (pokemon.Habitat, error) {
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

// GetShape returns a Shape struct containing information about the Shape with the given name.
func GetShape(name string) (pokemon.Shape, error) {
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

// GetSpecies returns a Species struct containing information about the Species with the given name.
func GetSpecies(name string) (pokemon.Species, error) {
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

// GetStat returns a Stat struct containing information about the Stat with the given name.
func GetStat(name string) (pokemon.Stat, error) {
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

// GetType returns a Type struct containing information about the Type with the given name.
func GetType(name string) (pokemon.Type, error) {
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
