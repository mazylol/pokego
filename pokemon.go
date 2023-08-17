package pokego

import (
	"encoding/json"
	"fmt"
	"github.com/mazylol/pokego/structs"
	"io"
	"net/http"
)

// GetPokemon returns a Pokemon struct containing information about the Pokemon with the given name.
func GetPokemon(name string) (structs.Pokemon, error) {
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
