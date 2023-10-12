package encounters

type Method struct {
	// The unique identifier for the method.
	Id int `json:"id"`

	// The name of the method.
	Name string `json:"name"`

	// The order associated with the method.
	Order int `json:"order"`

	// Names associated with the method in different languages.
	Names []struct {
		// The name of the method.
		Name string `json:"name"`

		// Information about the language of the method name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}
