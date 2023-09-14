package pokego

// Pokemon represents a response from the PokeAPI containing information about a specific Pokemon.
type Pokemon struct {
	// Abilities is an array containing information about the different abilities of the Pokémon.
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`

	// BaseExperience is the base experience gained when defeating or catching this Pokémon.
	BaseExperience int `json:"base_experience"`

	// Forms is an array containing information about the different forms of the Pokémon.
	Forms []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"forms"`

	// GameIndices is an array containing information about the game index of the Pokémon in different versions.
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`

	// Height is the height of the Pokémon in decimetres.
	Height int `json:"height"`

	// HeldItems is an array containing information about the items that the Pokémon can hold.
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"item"`

		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`

	// Id is the unique identifier of the Pokémon.
	Id int `json:"id"`

	// IsDefault indicates whether this is the default form of the Pokémon.
	IsDefault bool `json:"is_default"`

	// LocationAreaEncounters provides information about the areas where the Pokémon can be encountered.
	LocationAreaEncounters string `json:"location_area_encounters"`

	// Moves is an array containing information about the moves that the Pokémon can learn.
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"move"`

		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`

	// Name is the name of the Pokémon.
	Name string `json:"name"`

	// Order is the order in which the Pokémon appears in a Pokédex.
	Order int `json:"order"`

	// PastTypes is an array that appears to be reserved for future use, currently holding empty interfaces.
	PastTypes []interface{} `json:"past_types"`

	// Species provides information about the species of the Pokémon.
	Species struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"species"`

	// Sprites contains URLs to various sprite images representing the Pokémon.
	Sprites struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`

		// Other contains additional sprite information for various contexts.
		Other struct {
			DreamWorld struct {
				FrontDefault string      `json:"front_default"`
				FrontFemale  interface{} `json:"front_female"`
			} `json:"dream_world"`

			Home struct {
				FrontDefault     string      `json:"front_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShiny       string      `json:"front_shiny"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"home"`

			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
		} `json:"other"`

		Versions struct {
			// Detailed sprite information for various generations and versions.
			// The structure of this section isn't provided in the code, likely for brevity.
		} `json:"versions"`
	} `json:"sprites"`

	// Stats is an array containing information about the base stats of the Pokémon.
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`

	// Types is an array containing information about the types of the Pokémon.
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`

	// Weight is the weight of the Pokémon in hectograms.
	Weight int `json:"weight"`
}

// PokemonList represents a response from the PokeAPI containing a list of Pokémon.
type PokemonList struct {
	// Count is the total count of Pokémon in the list.
	Count int `json:"count"`

	// Next is the URL to the next page of the Pokémon list. It can be used for pagination.
	Next string `json:"next"`

	// Previous is the URL to the previous page of the Pokémon list. It can be used for pagination.
	// It is of type interface{} to handle cases where there is no previous page (null).
	Previous interface{} `json:"previous"`

	// Results is an array of struct entries containing information about individual Pokémon.
	Results []struct {
		// Name is the name of the Pokémon.
		Name string `json:"name"`

		// Url is the URL to the details of the specific Pokémon.
		Url string `json:"url"`
	} `json:"results"`
}

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
