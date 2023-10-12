package encounters

type Condition struct {
	// The unique identifier for the condition.
	Id int `json:"id"`

	// The name of the condition.
	Name string `json:"name"`

	// Values associated with the condition, each containing a name and URL.
	Values []struct {
		// The name of the condition value.
		Name string `json:"name"`

		// The URL to access detailed information about the condition value.
		Url string `json:"url"`
	} `json:"values"`

	// Names associated with the condition in different languages.
	Names []struct {
		// The name of the condition.
		Name string `json:"name"`

		// Information about the language of the condition name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}
