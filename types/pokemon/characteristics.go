package pokemon

type Characteristic struct {
	// The unique identifier for the characteristic.
	Id int `json:"id"`

	// The gene modulo associated with the characteristic.
	GeneModulo int `json:"gene_modulo"`

	// An array of possible values for the characteristic.
	PossibleValues []int `json:"possible_values"`

	// Information about the highest stat associated with the characteristic.
	HighestStat struct {
		// The name of the highest stat.
		Name string `json:"name"`

		// The URL to access detailed information about the highest stat.
		Url string `json:"url"`
	} `json:"highest_stat"`

	// Descriptions of the characteristic in different languages.
	Descriptions []struct {
		// The description of the characteristic.
		Description string `json:"description"`

		// Information about the language of the description.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
}
