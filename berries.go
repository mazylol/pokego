package pokego

import (
	"encoding/json"
	"fmt"
	"log"

	berriesCache "github.com/mazylol/pokego/cache/berries"
	"github.com/mazylol/pokego/types/berries"
)

// GetBerry returns a berry by name.
func GetBerry(name string) (berries.Berry, error) {
	berry, err := berriesCache.GetBerryFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("berry/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &berry)

		if err == nil {
			err = berriesCache.AddBerryToCache(berry)
		}

		return berry, err
	} else {
		return berry, err
	}
}

// GetBerryFirmness returns a berry firmness by name.
func GetBerryFirmness(name string) (berries.Firmness, error) {
	firmness, err := berriesCache.GetFirmnessFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("berry-firmness/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &firmness)

		if err == nil {
			err = berriesCache.AddFirmnessToCache(firmness)
		}

		return firmness, err
	} else {
		return firmness, err
	}
}

// GetBerryFlavor returns a berry flavor by name.
func GetBerryFlavor(name string) (berries.Flavor, error) {
	flavor, err := berriesCache.GetFlavorFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("berry-flavor/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &flavor)

		if err == nil {
			err = berriesCache.AddFlavorToCache(flavor)
		}

		return flavor, err
	} else {
		return flavor, err
	}
}
