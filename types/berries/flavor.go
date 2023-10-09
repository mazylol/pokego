package berries

type Flavor struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Berries []struct {
		Potency int `json:"potency"`
		Berry   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"berry"`
	} `json:"berries"`
	ContestType struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"contest_type"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}
