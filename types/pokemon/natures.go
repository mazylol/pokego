package pokemon

type Nature struct {
	// The unique identifier for the nature.
	Id int `json:"id"`

	// The name of the nature.
	Name string `json:"name"`

	// The decreased stat associated with the nature.
	DecreasedStat struct {
		// The name of the decreased stat.
		Name string `json:"name"`

		// The URL to access detailed information about the decreased stat.
		Url string `json:"url"`
	} `json:"decreased_stat"`

	// The increased stat associated with the nature.
	IncreasedStat struct {
		// The name of the increased stat.
		Name string `json:"name"`

		// The URL to access detailed information about the increased stat.
		Url string `json:"url"`
	} `json:"increased_stat"`

	// The flavor that the Pokémon likes when this nature is selected.
	LikesFlavor struct {
		// The name of the flavor.
		Name string `json:"name"`

		// The URL to access detailed information about the flavor.
		Url string `json:"url"`
	} `json:"likes_flavor"`

	// The flavor that the Pokémon dislikes when this nature is selected.
	HatesFlavor struct {
		// The name of the flavor.
		Name string `json:"name"`

		// The URL to access detailed information about the flavor.
		Url string `json:"url"`
	} `json:"hates_flavor"`

	// Changes to Pokeathlon stats associated with the nature.
	PokeathlonStatChanges []struct {
		// The maximum change in the stat.
		MaxChange int `json:"max_change"`

		// The Pokeathlon stat associated with the change.
		PokeathlonStat struct {
			// The name of the Pokeathlon stat.
			Name string `json:"name"`

			// The URL to access detailed information about the Pokeathlon stat.
			Url string `json:"url"`
		} `json:"pokeathlon_stat"`
	} `json:"pokeathlon_stat_changes"`

	// Preferences for move battle styles associated with the nature.
	MoveBattleStylePreferences []struct {
		// The preference for move battle style with low HP.
		LowHpPreference int `json:"low_hp_preference"`

		// The preference for move battle style with high HP.
		HighHpPreference int `json:"high_hp_preference"`

		// The move battle style associated with the preferences.
		MoveBattleStyle struct {
			// The name of the move battle style.
			Name string `json:"name"`

			// The URL to access detailed information about the move battle style.
			Url string `json:"url"`
		} `json:"move_battle_style"`
	} `json:"move_battle_style_preferences"`

	// Names of the nature in different languages.
	Names []struct {
		// The name of the nature in a specific language.
		Name string `json:"name"`

		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}
