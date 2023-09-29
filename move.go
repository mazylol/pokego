package pokego

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mazylol/pokego/cache"
	"github.com/mazylol/pokego/types"
)

// GetMove returns a Move struct containing information about the Move with the given name.
func GetMove(name string) (types.Move, error) {
	move, err := cache.GetMoveFromCache(name)
	if err != nil {
		body, err := callApi(fmt.Sprintf("move/%v", name))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &move)

		if err == nil {
			err = cache.AddMoveToCache(move)
		}

		return move, err
	} else {
		return move, err
	}
}

// GetMoveList returns a list of Move names. You have to include a limit for the amount of names you want.
func GetMoveList(limit int) (types.MoveList, error) {
	moveList, err := cache.GetMoveListFromCache(limit)
	if err != nil {
		body, err := callApi(fmt.Sprintf("move?limit=%v", limit))

		if err != nil {
			log.Fatal("Failed to call api")
		}

		err = json.Unmarshal(body, &moveList)

		if err == nil {
			err = cache.AddMoveListToCache(moveList, limit)
		}

		return moveList, err
	} else {
		return moveList, err
	}
}
