package moves

type BattleStyle struct {
	// The unique identifier for the battle style.
	ID int `json:"id"`

	// The name of the battle style.
	Name string `json:"name"`

	// Names of the battle style in different languages.
	Names []struct {
		// The name of the battle style in a specific language.
		Name string `json:"name"`

		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			URL string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}
