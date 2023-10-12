package contests

type Effect struct {
	// The unique identifier for the effect.
	Id int `json:"id"`

	// The appeal value of the effect.
	Appeal int `json:"appeal"`

	// The jam value of the effect.
	Jam int `json:"jam"`

	// Effect entries associated with the effect.
	EffectEntries []struct {
		// The description of the effect.
		Effect string `json:"effect"`

		// Information about the language of the effect description.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`

	// Flavor text entries associated with the effect.
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
}
