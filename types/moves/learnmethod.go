package moves

type LearnMethod struct {
	// The unique identifier for the learn method.
	ID int `json:"id"`

	// The name of the learn method.
	Name string `json:"name"`

	// Names of the learn method in different languages.
	Names []struct {
		// The name of the learn method.
		Name string `json:"name"`

		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			URL string `json:"url"`
		} `json:"language"`
	} `json:"names"`

	// Descriptions of the learn method in different languages.
	Descriptions []struct {
		// The description of the learn method.
		Description string `json:"description"`

		// Information about the language of the description.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			URL string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`

	// Version groups associated with this learn method.
	VersionGroups []struct {
		// The name of the version group.
		Name string `json:"name"`

		// The URL to access detailed information about the version group.
		URL string `json:"url"`
	} `json:"version_groups"`
}
