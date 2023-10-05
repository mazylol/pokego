package moves

type DamageClass struct {
	// The unique identifier for the damage class.
	ID int `json:"id"`

	// The name of the damage class.
	Name string `json:"name"`

	// Descriptions of the damage class in different languages.
	Descriptions []struct {
		// The description of the damage class.
		Description string `json:"description"`

		// Information about the language of the description.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			URL string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`

	// Moves associated with this damage class.
	Moves []struct {
		// The name of the move.
		Name string `json:"name"`

		// The URL to access detailed information about the move.
		URL string `json:"url"`
	} `json:"moves"`
}
