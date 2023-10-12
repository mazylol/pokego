package pokego

import (
	"encoding/json"
	"fmt"
	"log"

	encountersCache "github.com/mazylol/pokego/cache/encounters"
	"github.com/mazylol/pokego/types/encounters"
)

// GetEncounterMethod returns a EncounterMethod struct containing information about the EncounterMethod with the given name.
func GetEncounterMethod(name string) (encounters.Method, error) {
	method, err := encountersCache.GetMethodFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("encounter-method/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &method)

		if err == nil {
			err = encountersCache.AddMethodToCache(method)
		}

		return method, err
	} else {
		return method, err
	}
}

// GetEncounterCondition returns a EncounterCondition struct containing information about the EncounterCondition with the given name.
func GetEncounterCondition(name string) (encounters.Condition, error) {
	condition, err := encountersCache.GetConditionFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("encounter-condition/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &condition)

		if err == nil {
			err = encountersCache.AddConditionToCache(condition)
		}

		return condition, err
	} else {
		return condition, err
	}
}

// GetEncounterConditionValue returns a EncounterConditionValue struct containing information about the EncounterConditionValue with the given name.
func GetEncounterConditionValue(name string) (encounters.ConditionValue, error) {
	conditionValue, err := encountersCache.GetConditionValueFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("encounter-condition-value/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &conditionValue)

		if err == nil {
			err = encountersCache.AddConditionValueToCache(conditionValue)
		}

		return conditionValue, err
	} else {
		return conditionValue, err
	}
}
