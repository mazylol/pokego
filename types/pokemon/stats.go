package pokemon

type Stat struct {
	// Information about moves affecting this stat.
	AffectingMoves struct {
		// Moves that decrease this stat and their change value.
		Decrease []struct {
			// The change value for the stat decrease.
			Change int `json:"change"`

			// Information about the move causing the stat decrease.
			Move struct {
				// The name of the move.
				Name string `json:"name"`

				// The URL to access detailed information about the move.
				Url string `json:"url"`
			} `json:"move"`
		} `json:"decrease"`

		// Moves that increase this stat and their change value.
		Increase []struct {
			// The change value for the stat increase.
			Change int `json:"change"`

			// Information about the move causing the stat increase.
			Move struct {
				// The name of the move.
				Name string `json:"name"`

				// The URL to access detailed information about the move.
				Url string `json:"url"`
			} `json:"move"`
		} `json:"increase"`
	} `json:"affecting_moves"`

	// Information about natures affecting this stat.
	AffectingNatures struct {
		// Natures that decrease this stat.
		Decrease []struct {
			// The name of the nature.
			Name string `json:"name"`

			// The URL to access detailed information about the nature.
			Url string `json:"url"`
		} `json:"decrease"`

		// Natures that increase this stat.
		Increase []struct {
			// The name of the nature.
			Name string `json:"name"`

			// The URL to access detailed information about the nature.
			Url string `json:"url"`
		} `json:"increase"`
	} `json:"affecting_natures"`

	// Characteristics associated with this stat.
	Characteristics []struct {
		// The URL to access detailed information about the characteristic.
		Url string `json:"url"`
	} `json:"characteristics"`

	// The game index of this stat.
	GameIndex int `json:"game_index"`

	// The unique identifier for the stat.
	Id int `json:"id"`

	// Indicates whether this stat is for battles only.
	IsBattleOnly bool `json:"is_battle_only"`

	// The damage class associated with moves affecting this stat.
	MoveDamageClass struct {
		// The name of the move damage class.
		Name string `json:"name"`

		// The URL to access detailed information about the move damage class.
		Url string `json:"url"`
	} `json:"move_damage_class"`

	// The name of the stat.
	Name string `json:"name"`

	// Names of the stat in different languages.
	Names []struct {
		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`

		// The name of the stat in a specific language.
		Name string `json:"name"`
	} `json:"names"`
}
