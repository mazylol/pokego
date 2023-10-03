package pokemon

type Ability struct {
	// The unique identifier for the ability.
	Id int `json:"id"`

	// The name of the ability.
	Name string `json:"name"`

	// Indicates whether the ability is part of the main series Pokémon games.
	IsMainSeries bool `json:"is_main_series"`

	// Information about the generation in which the ability was introduced.
	Generation struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"generation"`

	// Names of the ability in different languages.
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`

	// Descriptions of the ability's effects in different languages.
	EffectEntries []struct {
		Effect      string `json:"effect"`
		ShortEffect string `json:"short_effect"`
		Language    struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`

	// Changes to the ability's effects in different version groups.
	EffectChanges []struct {
		VersionGroup struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version_group"`
		EffectEntries []struct {
			Effect   string `json:"effect"`
			Language struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"language"`
		} `json:"effect_entries"`
	} `json:"effect_changes"`

	// Descriptions of the ability's flavor text in different languages and version groups.
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
		VersionGroup struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`

	// Pokémon that have this ability.
	Pokemon []struct {
		// Indicates if the ability is hidden (only available in certain conditions).
		IsHidden bool `json:"is_hidden"`

		// The slot in which the ability is assigned for the Pokémon.
		Slot int `json:"slot"`

		// Information about the Pokémon with this ability.
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon"`
}
