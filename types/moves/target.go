package moves

type Target struct {
	// The unique identifier for the target.
	ID int `json:"id"`

	// The name of the target.
	Name string `json:"name"`

	// Descriptions of the target in different languages.
	Descriptions []struct {
		// The description of the target.
		Description string `json:"description"`

		// Information about the language of the description.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			URL string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`

	// Moves associated with this target.
	Moves []struct {
		// The name of the move.
		Name string `json:"name"`

		// The URL to access detailed information about the move.
		URL string `json:"url"`
	} `json:"moves"`

	// Names of the target in different languages.
	Names []struct {
		// The name of the target.
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
