package pokemon

type Color struct {
	// The unique identifier for the color.
	Id int `json:"id"`

	// The name of the color.
	Name string `json:"name"`

	// Names of the color in different languages.
	Names []struct {
		// The name of the color in a specific language.
		Name string `json:"name"`

		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"names"`

	// Pokémon species that have this color.
	PokemonSpecies []struct {
		// The name of the Pokémon species.
		Name string `json:"name"`

		// The URL to access detailed information about the Pokémon species.
		Url string `json:"url"`
	} `json:"pokemon_species"`
}
