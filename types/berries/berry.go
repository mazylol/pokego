package berries

type Berry struct {
	// The unique identifier for the berry.
	ID int `json:"id"`

	// The name of the berry.
	Name string `json:"name"`

	// The time it takes for the berry to grow.
	GrowthTime int `json:"growth_time"`

	// The maximum number of times the berry can be harvested from a plant.
	MaxHarvest int `json:"max_harvest"`

	// The power of the Natural Gift move when used with this berry.
	NaturalGiftPower int `json:"natural_gift_power"`

	// The size of the berry.
	Size int `json:"size"`

	// The smoothness of the berry's skin.
	Smoothness int `json:"smoothness"`

	// The dryness level of the soil in which the berry is planted.
	SoilDryness int `json:"soil_dryness"`

	// The firmness of the berry.
	Firmness struct {
		// The name of the firmness.
		Name string `json:"name"`

		// The URL to access detailed information about the firmness.
		URL string `json:"url"`
	} `json:"firmness"`

	// Flavors associated with the berry.
	Flavors []struct {
		// The potency of the flavor.
		Potency int `json:"potency"`

		// Information about the flavor.
		Flavor struct {
			// The name of the flavor.
			Name string `json:"name"`

			// The URL to access detailed information about the flavor.
			URL string `json:"url"`
		} `json:"flavor"`
	} `json:"flavors"`

	// The item form of the berry.
	Item struct {
		// The name of the item.
		Name string `json:"name"`

		// The URL to access detailed information about the item.
		URL string `json:"url"`
	} `json:"item"`

	// The type of move that Natural Gift becomes when used with this berry.
	NaturalGiftType struct {
		// The name of the type.
		Name string `json:"name"`

		// The URL to access detailed information about the type.
		URL string `json:"url"`
	} `json:"natural_gift_type"`
}
