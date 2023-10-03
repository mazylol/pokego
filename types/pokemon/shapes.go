package pokemon

type Shape struct {
	// Awesome names associated with the shape.
	AwesomeNames []struct {
		// The awesome name for the shape.
		AwesomeName string `json:"awesome_name"`

		// Information about the language of the awesome name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"awesome_names"`

	// The unique identifier for the shape.
	Id int `json:"id"`

	// The name of the shape.
	Name string `json:"name"`

	// Names of the shape in different languages.
	Names []struct {
		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`

		// The name of the shape in a specific language.
		Name string `json:"name"`
	} `json:"names"`

	// Pokémon species that have this shape.
	PokemonSpecies []struct {
		// The name of the Pokémon species.
		Name string `json:"name"`

		// The URL to access detailed information about the Pokémon species.
		Url string `json:"url"`
	} `json:"pokemon_species"`
}
