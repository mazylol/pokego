package structs

// PokemonList represents a response from the PokeAPI containing a list of Pokémon.
type PokemonList struct {
	// Count is the total count of Pokémon in the list.
	Count int `json:"count"`

	// Next is the URL to the next page of the Pokémon list. It can be used for pagination.
	Next string `json:"next"`

	// Previous is the URL to the previous page of the Pokémon list. It can be used for pagination.
	// It is of type interface{} to handle cases where there is no previous page (null).
	Previous interface{} `json:"previous"`

	// Results is an array of struct entries containing information about individual Pokémon.
	Results []struct {
		// Name is the name of the Pokémon.
		Name string `json:"name"`

		// Url is the URL to the details of the specific Pokémon.
		Url string `json:"url"`
	} `json:"results"`
}
