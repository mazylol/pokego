package pokego

import (
	"encoding/json"
	"fmt"
	movesCache "github.com/mazylol/pokego/cache/moves"
	"github.com/mazylol/pokego/types/moves"
	"log"
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
