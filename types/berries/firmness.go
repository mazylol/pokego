package berries

type Firmness struct {
	// The unique identifier for the firmness.
	ID int `json:"id"`

	// The name of the firmness.
	Name string `json:"name"`

	// Berries associated with this firmness.
	Berries []struct {
		// The name of the berry.
		Name string `json:"name"`

		// The URL to access detailed information about the berry.
		URL string `json:"url"`
	} `json:"berries"`

	// Names of the firmness in different languages.
	Names []struct {
		// The name of the firmness in a specific language.
		Name string `json:"name"`

		// Information about the language.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			URL string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}
