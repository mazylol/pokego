package pokemon

type Habitat struct {
	// The unique identifier for the habitat.
	Id int `json:"id"`

	// The name of the habitat.
	Name string `json:"name"`

	// Names of the habitat in different languages.
	Names []struct {
		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`

		// The name of the habitat in a specific language.
		Name string `json:"name"`
	} `json:"names"`

	// Pokémon species that can be found in this habitat.
	PokemonSpecies []struct {
		// The name of the Pokémon species.
		Name string `json:"name"`

		// The URL to access detailed information about the Pokémon species.
		Url string `json:"url"`
	} `json:"pokemon_species"`
}
