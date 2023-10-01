package pokemon

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
