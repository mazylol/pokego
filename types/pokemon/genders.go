package pokemon

type Gender struct {
	// The unique identifier for the gender.
	Id int `json:"id"`

	// The name of the gender (e.g., "Male," "Female," "Genderless").
	Name string `json:"name"`

	// Details about Pokémon species and their gender rates.
	PokemonSpeciesDetails []struct {
		// The rate or percentage of this gender in the population.
		Rate int `json:"rate"`

		// Information about the Pokémon species.
		PokemonSpecies struct {
			// The name of the Pokémon species.
			Name string `json:"name"`

			// The URL to access detailed information about the Pokémon species.
			Url string `json:"url"`
		} `json:"pokemon_species"`
	} `json:"pokemon_species_details"`

	// List of Pokémon species for which this gender is required for evolution.
	RequiredForEvolution []struct {
		// The name of the Pokémon species.
		Name string `json:"name"`

		// The URL to access detailed information about the Pokémon species.
		Url string `json:"url"`
	} `json:"required_for_evolution"`
}
