package pokemon

type EggGroup struct {
	// The unique identifier for the Egg Group.
	Id int `json:"id"`

	// The name of the Egg Group.
	Name string `json:"name"`

	// Names of the Egg Group in different languages.
	Names []struct {
		// The name of the Egg Group in a specific language.
		Name string `json:"name"`

		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"names"`

	// Pokémon species that belong to this Egg Group.
	PokemonSpecies []struct {
		// The name of the Pokémon species.
		Name string `json:"name"`

		// The URL to access detailed information about the Pokémon species.
		Url string `json:"url"`
	} `json:"pokemon_species"`
}
