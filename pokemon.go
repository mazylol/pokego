package pokego

import (
	"encoding/json"
	"fmt"
	"github.com/mazylol/pokego/structs"
	"io"
	"net/http"
)

type idType interface {
	int | string
}

// GetPokemon returns a Pokemon struct containing information about the Pokémon with the given id/name.
func GetPokemon[identifier idType](id identifier) (structs.Pokemon, error) {
	var url = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", id)

	req, err := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)

	var pokemon structs.Pokemon
	err = json.Unmarshal(body, &pokemon)

	return pokemon, err
}
