package pokego

import (
	"encoding/json"
	"fmt"
	"log"

	movesCache "github.com/mazylol/pokego/cache/moves"
	"github.com/mazylol/pokego/types/moves"
)

// GetMove returns a Move struct containing information about the Move with the given name.
func GetMove(name string) (moves.Move, error) {
	mov, err := movesCache.GetMoveFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("move/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &mov)

		if err == nil {
			err = movesCache.AddMoveToCache(mov)
		}

		return mov, err
	} else {
		return mov, err
	}
}

// GetAilment returns a Move struct containing information about the Move with the given name.
func GetAilment(name string) (moves.Ailment, error) {
	ailment, err := movesCache.GetAilmentFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("move-ailment/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &ailment)

		if err == nil {
			err = movesCache.AddAilmentToCache(ailment)
		}

		return ailment, err
	} else {
		return ailment, err
	}
}

// GetBattleStyle returns a Move struct containing information about the Move with the given name.
func GetBattleStyle(name string) (moves.BattleStyle, error) {
	battleStyle, err := movesCache.GetBattleStyleFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("move-battle-style/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &battleStyle)

		if err == nil {
			err = movesCache.AddBattleStyleToCache(battleStyle)
		}

		return battleStyle, err
	} else {
		return battleStyle, err
	}
}

// GetCategory returns a Move struct containing information about the Move with the given name.
func GetCategory(name string) (moves.Category, error) {
	category, err := movesCache.GetCategoryFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("move-category/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &category)

		if err == nil {
			err = movesCache.AddCategoryToCache(category)
		}

		return category, err
	} else {
		return category, err
	}
}

// GetDamageClass returns a Move struct containing information about the Move with the given name.
func GetDamageClass(name string) (moves.DamageClass, error) {
	damageClass, err := movesCache.GetDamageClassFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("move-damage-class/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &damageClass)

		if err == nil {
			err = movesCache.AddDamageClassToCache(damageClass)
		}

		return damageClass, err
	} else {
		return damageClass, err
	}
}

// GetLearnMethod returns a Move struct containing information about the Move with the given name.
func GetLearnMethod(name string) (moves.LearnMethod, error) {
	learnMethod, err := movesCache.GetLearnMethodFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("move-learn-method/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &learnMethod)

		if err == nil {
			err = movesCache.AddLearnMethodToCache(learnMethod)
		}

		return learnMethod, err
	} else {
		return learnMethod, err
	}
}

// GetTarget returns a Move struct containing information about the Move with the given name.
func GetTarget(name string) (moves.Target, error) {
	target, err := movesCache.GetTargetFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("move-target/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &target)

		if err == nil {
			err = movesCache.AddTargetToCache(target)
		}

		return target, err
	} else {
		return target, err
	}
}
