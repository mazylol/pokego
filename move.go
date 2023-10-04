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

// GetMove returns a Move struct containing information about the Move with the given name.
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
