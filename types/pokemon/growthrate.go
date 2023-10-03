package pokemon

type GrowthRate struct {
	// The unique identifier for the growth rate.
	Id int `json:"id"`

	// The name of the growth rate.
	Name string `json:"name"`

	// The formula used to calculate the growth rate.
	Formula string `json:"formula"`

	// Descriptions of the growth rate in different languages.
	Descriptions []struct {
		// The description of the growth rate.
		Description string `json:"description"`

		// Information about the language of the description.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`

	// List of levels and corresponding experience points required for growth.
	Levels []struct {
		// The level at which the experience points are required.
		Level int `json:"level"`

		// The experience points required to reach the specified level.
		Experience int `json:"experience"`
	} `json:"levels"`

	// Pokémon species that share this growth rate.
	PokemonSpecies []struct {
		// The name of the Pokémon species.
		Name string `json:"name"`

		// The URL to access detailed information about the Pokémon species.
		Url string `json:"url"`
	} `json:"pokemon_species"`
}
