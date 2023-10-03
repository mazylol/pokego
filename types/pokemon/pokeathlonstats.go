package pokemon

type PokeathlonStat struct {
	// The unique identifier for the Pokeathlon Stat.
	Id int `json:"id"`

	// The name of the Pokeathlon Stat.
	Name string `json:"name"`

	// Affecting natures that can increase or decrease this stat.
	AffectingNatures struct {
		// Natures that increase this stat and their maximum change values.
		Increase []struct {
			// The maximum change in the stat.
			MaxChange int `json:"max_change"`

			// The nature associated with the increase.
			Nature struct {
				// The name of the nature.
				Name string `json:"name"`

				// The URL to access detailed information about the nature.
				Url string `json:"url"`
			} `json:"nature"`
		} `json:"increase"`

		// Natures that decrease this stat and their maximum change values.
		Decrease []struct {
			// The maximum change in the stat.
			MaxChange int `json:"max_change"`

			// The nature associated with the decrease.
			Nature struct {
				// The name of the nature.
				Name string `json:"name"`

				// The URL to access detailed information about the nature.
				Url string `json:"url"`
			} `json:"nature"`
		} `json:"decrease"`
	} `json:"affecting_natures"`

	// Names of the Pokeathlon Stat in different languages.
	Names []struct {
		// The name of the Pokeathlon Stat in a specific language.
		Name string `json:"name"`

		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}
