package contests

type Type struct {
	// The unique identifier for the type.
	Id int `json:"id"`

	// The name of the type.
	Name string `json:"name"`

	// Information about the berry flavor associated with the type.
	BerryFlavor struct {
		// The name of the berry flavor.
		Name string `json:"name"`

		// The URL to access detailed information about the berry flavor.
		Url string `json:"url"`
	} `json:"berry_flavor"`

	// Names associated with the type in different languages.
	Names []struct {
		// The name of the type.
		Name string `json:"name"`

		// The color associated with the type.
		Color string `json:"color"`

		// Information about the language of the type name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}
