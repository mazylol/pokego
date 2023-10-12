package pokego

import (
	"encoding/json"
	"fmt"
	"log"

	contestsCache "github.com/mazylol/pokego/cache/contests"
	"github.com/mazylol/pokego/types/contests"
)

// GetContestType returns a ContestType struct containing information about the ContestType with the given name.
func GetContestType(name string) (contests.Type, error) {
	tipe, err := contestsCache.GetTypeFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("contest-type/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &tipe)

		if err == nil {
			err = contestsCache.AddTypeToCache(tipe)
		}

		return tipe, err
	} else {
		return tipe, err
	}
}

// GetContestEffect returns a ContestEffect struct containing information about the ContestEffect with the given id.
func GetContestEffect(id int) (contests.Effect, error) {
	effect, err := contestsCache.GetEffectFromCache(id)
	if err != nil {
		body, err := callApi(fmt.Sprintf("contest-effect/%v", id))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &effect)

		if err == nil {
			err = contestsCache.AddEffectToCache(effect)
		}

		return effect, err
	} else {
		return effect, err
	}
}

// GetSuperContestEffect returns a SuperContestEffect struct containing information about the SuperContestEffect with the given id.
func GetSuperContestEffect(id int) (contests.SuperEffect, error) {
	superEffect, err := contestsCache.GetSuperEffectFromCache(id)
	if err != nil {
		body, err := callApi(fmt.Sprintf("super-contest-effect/%v", id))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &superEffect)

		if err == nil {
			err = contestsCache.AddSuperEffectToCache(superEffect)
		}

		return superEffect, err
	} else {
		return superEffect, err
	}
}
