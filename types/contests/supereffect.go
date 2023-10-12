package contests

type SuperEffect struct {
	// The unique identifier for the super effect.
	Id int `json:"id"`

	// The appeal of the super effect.
	Appeal int `json:"appeal"`

	// Flavor text entries associated with the super effect.
	FlavorTextEntries []struct {
		// The flavor text description.
		FlavorText string `json:"flavor_text"`

		// Information about the language of the flavor text.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"flavor_text_entries"`

	// Moves associated with this super effect.
	Moves []struct {
		// The name of the move.
		Name string `json:"name"`

		// The URL to access detailed information about the move.
		Url string `json:"url"`
	} `json:"moves"`
}
