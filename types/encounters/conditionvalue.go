package encounters

type ConditionValue struct {
	// Information about the condition associated with this value.
	Condition struct {
		// The name of the condition.
		Name string `json:"name"`

		// The URL to access detailed information about the condition.
		Url string `json:"url"`
	} `json:"condition"`

	// The unique identifier for the condition value.
	Id int `json:"id"`

	// The name of the condition value.
	Name string `json:"name"`

	// Names associated with the condition value in different languages.
	Names []struct {
		// The name of the condition value.
		Name string `json:"name"`

		// Information about the language of the condition value name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}
