package berries

type Flavor struct {
	// The unique identifier for the flavor.
	ID int `json:"id"`

	// The name of the flavor.
	Name string `json:"name"`

	// Berries associated with this flavor.
	Berries []struct {
		// The potency of the flavor in berries.
		Potency int `json:"potency"`

		// Information about the associated berry.
		Berry struct {
			// The name of the berry.
			Name string `json:"name"`

			// The URL to access detailed information about the berry.
			URL string `json:"url"`
		} `json:"berry"`
	} `json:"berries"`

	// Contest type associated with this flavor.
	ContestType struct {
		// The name of the contest type.
		Name string `json:"name"`

		// The URL to access detailed information about the contest type.
		URL string `json:"url"`
	} `json:"contest_type"`

	// Names of the flavor in different languages.
	Names []struct {
		// The name of the flavor in a specific language.
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
