package moves

// Move represents a response from the PokeAPI containing information about a specific Move.
type Move struct {
	// The unique identifier for the move.
	Id int `json:"id"`

	// The name of the move.
	Name string `json:"name"`

	// The move's accuracy.
	Accuracy int `json:"accuracy"`

	// The chance of an additional effect occurring when using the move.
	EffectChance interface{} `json:"effect_chance"`

	// The power points (PP) of the move, representing the number of times it can be used.
	Pp int `json:"pp"`

	// The move's priority in battle.
	Priority int `json:"priority"`

	// The power of the move, indicating its strength.
	Power int `json:"power"`

	// Information about contest combos for the move.
	ContestCombos struct {
		Normal struct {
			// Moves that can be used before this move in a contest combo.
			UseBefore []struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"use_before"`

			// A move that can be used after this move in a contest combo.
			UseAfter interface{} `json:"use_after"`
		} `json:"normal"`

		Super struct {
			// Moves that can be used before this move in a super contest combo.
			UseBefore interface{} `json:"use_before"`

			// A move that can be used after this move in a super contest combo.
			UseAfter interface{} `json:"use_after"`
		} `json:"super"`
	} `json:"contest_combos"`

	// The contest type associated with the move.
	ContestType struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"contest_type"`

	// The effect of the move in a contest.
	ContestEffect struct {
		Url string `json:"url"`
	} `json:"contest_effect"`

	// The damage class of the move (e.g., physical, special).
	DamageClass struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"damage_class"`

	// Descriptions of the move's effects in various languages.
	EffectEntries []struct {
		Effect      string `json:"effect"`
		ShortEffect string `json:"short_effect"`
		Language    struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`

	// Changes to the move's effects.
	EffectChanges []interface{} `json:"effect_changes"`

	// The generation in which the move was introduced.
	Generation struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"generation"`

	// Metadata about the move's properties.
	Meta struct {
		Ailment struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"ailment"`

		Category struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"category"`

		// Minimum and maximum number of hits for the move.
		MinHits interface{} `json:"min_hits"`
		MaxHits interface{} `json:"max_hits"`

		// Minimum and maximum number of turns for the move.
		MinTurns interface{} `json:"min_turns"`
		MaxTurns interface{} `json:"max_turns"`

		// Amount of HP drained from the opponent.
		Drain int `json:"drain"`

		// Amount of HP healed by the move.
		Healing int `json:"healing"`

		// Critical hit rate.
		CritRate int `json:"crit_rate"`

		// Chance of causing an ailment.
		AilmentChance int `json:"ailment_chance"`

		// Chance of causing the opponent to flinch.
		FlinchChance int `json:"flinch_chance"`

		// Chance of altering opponent's stats.
		StatChance int `json:"stat_chance"`
	} `json:"meta"`

	// Names of the move in different languages.
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`

	// Previous values of the move (if any).
	PastValues []interface{} `json:"past_values"`

	// Changes to the move's statistics.
	StatChanges []interface{} `json:"stat_changes"`

	// Information about the move's super contest effect.
	SuperContestEffect struct {
		Url string `json:"url"`
	} `json:"super_contest_effect"`

	// The target of the move (e.g., opponent, self).
	Target struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"target"`

	// The type of the move (e.g., Water, Fire).
	Type struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"type"`

	// List of Pokémon that can learn this move.
	LearnedByPokemon []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"learned_by_pokemon"`

	// Descriptions of the move's flavor text in various languages.
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Url  string `json:"url"`
			Name string `json:"name"`
		} `json:"language"`
		VersionGroup struct {
			Url  string `json:"url"`
			Name string `json:"name"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
}

// MoveList represents a response from the PokeAPI containing a list of Moves.
type MoveList struct {
	// The total count of moves in the list.
	Count int `json:"count"`

	// URL to the next page of moves (if available).
	Next string `json:"next"`

	// URL to the previous page of moves (if available).
	Previous interface{} `json:"previous"`

	// An array of move names and their corresponding URLs.
	Results []struct {
		// The name of the move.
		Name string `json:"name"`

		// The URL to access detailed information about the move.
		Url string `json:"url"`
	} `json:"results"`
}
