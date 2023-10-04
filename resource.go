package pokego

import (
	"encoding/json"
	"log"

	resourceCache "github.com/mazylol/pokego/cache/resource"
	"github.com/mazylol/pokego/types/resource"
)

// GetResourceList returns a ResourceList struct containing information about the list of resources with the given name.
func GetResourceList(endpoint string, count int) (resource.List, error) {
	resourceList, err := resourceCache.GetResourceListFromCache(endpoint, count)
	if err != nil {
		body, err := callApi(endpoint)

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &resourceList)

		if err == nil {
			err = resourceCache.AddResourceListToCache(resourceList, endpoint, count)
		}

		return resourceList, err
	} else {
		return resourceList, err
	}
}
