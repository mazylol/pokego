package pokemon

type Form struct {
	// The name of the Pokémon form.
	FormName string `json:"form_name"`

	// Names of the Pokémon form in different languages.
	FormNames []struct {
		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`

		// The name of the Pokémon form in a specific language.
		Name string `json:"name"`
	} `json:"form_names"`

	// The form order of the Pokémon.
	FormOrder int `json:"form_order"`

	// The unique identifier for the Pokémon form.
	Id int `json:"id"`

	// Indicates whether this form is for battles only.
	IsBattleOnly bool `json:"is_battle_only"`

	// Indicates whether this form is the default form.
	IsDefault bool `json:"is_default"`

	// Indicates whether this form is a Mega Evolution.
	IsMega bool `json:"is_mega"`

	// The name of the Pokémon associated with this form.
	Name string `json:"name"`

	// Names of the Pokémon in different languages.
	Names []struct {
		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`

		// The name of the Pokémon in a specific language.
		Name string `json:"name"`
	} `json:"names"`

	// The order of the Pokémon form.
	Order int `json:"order"`

	// Information about the Pokémon associated with this form.
	Pokemon struct {
		// The name of the Pokémon.
		Name string `json:"name"`

		// The URL to access detailed information about the Pokémon.
		Url string `json:"url"`
	} `json:"pokemon"`

	// Sprite images for the Pokémon form, including front and back views in different variations.
	Sprites struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"sprites"`

	// Types associated with the Pokémon form, including the slot and type information.
	Types []struct {
		Slot int `json:"slot"`

		Type struct {
			// The name of the Pokémon type.
			Name string `json:"name"`

			// The URL to access detailed information about the Pokémon type.
			Url string `json:"url"`
		} `json:"type"`
	} `json:"types"`

	// The version group in which this form appears.
	VersionGroup struct {
		// The name of the version group.
		Name string `json:"name"`

		// The URL to access detailed information about the version group.
		Url string `json:"url"`
	} `json:"version_group"`
}
