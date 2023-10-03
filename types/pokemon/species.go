package pokemon

type Species struct {
	BaseHappiness int `json:"base_happiness"`
	CaptureRate   int `json:"capture_rate"`
	Color         struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"color"`
	EggGroups []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"egg_groups"`
	EvolutionChain struct {
		Url string `json:"url"`
	} `json:"evolution_chain"`
	EvolvesFromSpecies struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"evolves_from_species"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
		Version struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version"`
	} `json:"flavor_text_entries"`
	FormDescriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"form_descriptions"`
	FormsSwitchable bool `json:"forms_switchable"`
	GenderRate      int  `json:"gender_rate"`
	Genera          []struct {
		Genus    string `json:"genus"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"genera"`
	Generation struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"generation"`
	GrowthRate struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"growth_rate"`
	Habitat              interface{} `json:"habitat"`
	HasGenderDifferences bool        `json:"has_gender_differences"`
	HatchCounter         int         `json:"hatch_counter"`
	Id                   int         `json:"id"`
	IsBaby               bool        `json:"is_baby"`
	IsLegendary          bool        `json:"is_legendary"`
	IsMythical           bool        `json:"is_mythical"`
	Name                 string      `json:"name"`
	Names                []struct {
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Order             int `json:"order"`
	PalParkEncounters []struct {
		Area struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"area"`
		BaseScore int `json:"base_score"`
		Rate      int `json:"rate"`
	} `json:"pal_park_encounters"`
	PokedexNumbers []struct {
		EntryNumber int `json:"entry_number"`
		Pokedex     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokedex"`
	} `json:"pokedex_numbers"`
	Shape struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"shape"`
	Varieties []struct {
		IsDefault bool `json:"is_default"`
		Pokemon   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"varieties"`
}
