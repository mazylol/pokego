package pokemon

type LocationArea struct {
	// Information about the location area.
	LocationArea struct {
		// The name of the location area.
		Name string `json:"name"`

		// The URL to access detailed information about the location area.
		Url string `json:"url"`
	} `json:"location_area"`

	// Details of encounters in different versions of the game.
	VersionDetails []struct {
		// The maximum chance for an encounter in this version.
		MaxChance int `json:"max_chance"`

		// Details of encounters in this version, including level range, conditions, chance, and method.
		EncounterDetails []struct {
			// The minimum level for the encounter.
			MinLevel int `json:"min_level"`

			// The maximum level for the encounter.
			MaxLevel int `json:"max_level"`

			// Condition values that may apply to the encounter.
			ConditionValues []struct {
				// The name of the condition value.
				Name string `json:"name"`

				// The URL to access detailed information about the condition value.
				Url string `json:"url"`
			} `json:"condition_values"`

			// The chance of encountering Pokémon with these conditions and levels.
			Chance int `json:"chance"`

			// The method of encounter.
			Method struct {
				// The name of the encounter method.
				Name string `json:"name"`

				// The URL to access detailed information about the encounter method.
				Url string `json:"url"`
			} `json:"method"`
		} `json:"encounter_details"`

		// Information about the version of the game.
		Version struct {
			// The name of the version.
			Name string `json:"name"`

			// The URL to access detailed information about the version.
			Url string `json:"url"`
		} `json:"version"`
	} `json:"version_details"`
}
