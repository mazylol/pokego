package moves

type Ailment struct {
	// The unique identifier for the ailment.
	ID int `json:"id"`

	// Moves associated with this ailment.
	Moves []struct {
		// The name of the move.
		Name string `json:"name"`

		// The URL to access detailed information about the move.
		URL string `json:"url"`
	} `json:"moves"`

	// The name of the ailment.
	Name string `json:"name"`

	// Names of the ailment in different languages.
	Names []struct {
		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			URL string `json:"url"`
		} `json:"language"`

		// The name of the ailment in a specific language.
		Name string `json:"name"`
	} `json:"names"`
}
