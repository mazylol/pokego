package pokego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetMove returns a Move struct containing information about the Move with the given name.
func GetMove(name string) (Move, error) {
	move, err := getMoveFromCache(name)
	if err != nil {
		var url = fmt.Sprintf("https://pokeapi.co/api/v2/move/%v", name)

		req, err := http.NewRequest("GET", url, nil)

		res, err := http.DefaultClient.Do(req)

		defer func(Body io.ReadCloser) {
			err = Body.Close()
		}(res.Body)

		body, err := io.ReadAll(res.Body)

		err = json.Unmarshal(body, &move)

		if err == nil {
			err = addMoveToCache(move)
		}

		return move, err
	} else {
		return move, err
	}
}

// GetMoveList returns a list of Move names. You have to include a limit for the amount of names you want.
func GetMoveList(limit int) (MoveList, error) {
	moveList, err := getMoveListFromCache(limit)
	if err != nil {
		var url = fmt.Sprintf("https://pokeapi.co/api/v2/move?limit=%v", limit)

		req, err := http.NewRequest("GET", url, nil)

		res, err := http.DefaultClient.Do(req)

		defer func(Body io.ReadCloser) {
			err = Body.Close()
		}(res.Body)

		body, err := io.ReadAll(res.Body)

		err = json.Unmarshal(body, &moveList)

		if err == nil {
			err = addMoveListToCache(moveList, limit)
		}

		return moveList, err
	} else {
		return moveList, err
	}
}
