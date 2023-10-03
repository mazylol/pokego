package pokemon

type Type struct {
	// Damage relations for this type.
	DamageRelations struct {
		// Types that deal double damage to this type.
		DoubleDamageFrom []struct {
			// The name of the type.
			Name string `json:"name"`

			// The URL to access detailed information about the type.
			Url string `json:"url"`
		} `json:"double_damage_from"`

		// Types that receive double damage from this type.
		DoubleDamageTo []struct {
			// The name of the type.
			Name string `json:"name"`

			// The URL to access detailed information about the type.
			Url string `json:"url"`
		} `json:"double_damage_to"`

		// Types that deal half damage to this type.
		HalfDamageFrom []struct {
			// The name of the type.
			Name string `json:"name"`

			// The URL to access detailed information about the type.
			Url string `json:"url"`
		} `json:"half_damage_from"`

		// Types that receive half damage from this type.
		HalfDamageTo []struct {
			// The name of the type.
			Name string `json:"name"`

			// The URL to access detailed information about the type.
			Url string `json:"url"`
		} `json:"half_damage_to"`

		// Types that deal no damage to this type.
		NoDamageFrom []struct {
			// The name of the type.
			Name string `json:"name"`

			// The URL to access detailed information about the type.
			Url string `json:"url"`
		} `json:"no_damage_from"`

		// Types that receive no damage from this type.
		NoDamageTo []struct {
			// The name of the type.
			Name string `json:"name"`

			// The URL to access detailed information about the type.
			Url string `json:"url"`
		} `json:"no_damage_to"`
	} `json:"damage_relations"`

	// Game indices associated with this type.
	GameIndices []struct {
		// The game index of this type.
		GameIndex int `json:"game_index"`

		// Information about the generation of the game.
		Generation struct {
			// The name of the generation.
			Name string `json:"name"`

			// The URL to access detailed information about the generation.
			Url string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`

	// Information about the generation this type belongs to.
	Generation struct {
		// The name of the generation.
		Name string `json:"name"`

		// The URL to access detailed information about the generation.
		Url string `json:"url"`
	} `json:"generation"`

	// The unique identifier for the type.
	Id int `json:"id"`

	// The damage class associated with moves of this type.
	MoveDamageClass struct {
		// The name of the move damage class.
		Name string `json:"name"`

		// The URL to access detailed information about the move damage class.
		Url string `json:"url"`
	} `json:"move_damage_class"`

	// Moves that belong to this type.
	Moves []struct {
		// The name of the move.
		Name string `json:"name"`

		// The URL to access detailed information about the move.
		Url string `json:"url"`
	} `json:"moves"`

	// The name of the type.
	Name string `json:"name"`

	// Names of the type in different languages.
	Names []struct {
		// Information about the language of the name.
		Language struct {
			// The name of the language.
			Name string `json:"name"`

			// The URL to access detailed information about the language.
			Url string `json:"url"`
		} `json:"language"`

		// The name of the type in a specific language.
		Name string `json:"name"`
	} `json:"names"`

	// Past damage relations for this type (deprecated).
	PastDamageRelations []interface{} `json:"past_damage_relations"`

	// Pokémon that have this type and their associated slot.
	Pokemon []struct {
		// Information about the Pokémon with this type.
		Pokemon struct {
			// The name of the Pokémon.
			Name string `json:"name"`

			// The URL to access detailed information about the Pokémon.
			Url string `json:"url"`
		} `json:"pokemon"`

		// The slot associated with the Pokémon.
		Slot int `json:"slot"`
	} `json:"pokemon"`
}
