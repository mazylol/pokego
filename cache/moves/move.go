package moves

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/moves"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddMoveToCache(move moves.Move) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&move)

	_, err = db.Exec("INSERT INTO move (name, data) VALUES (?, ?)", move.Name, string(jayson))

	return err
}

func GetMoveFromCache(name string) (moves.Move, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var move moves.Move

	row := db.QueryRow("SELECT data FROM move WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &move)

	return move, err
}

func AddMoveListToCache(moveList moves.MoveList, count int) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&moveList)

	_, err = db.Exec("INSERT INTO move_list (count, data) VALUES (?, ?)", count, string(jayson))

	return err
}

func GetMoveListFromCache(count int) (moves.MoveList, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var moveList moves.MoveList

	row := db.QueryRow("SELECT data FROM move_list WHERE count = ?", count)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &moveList)

	return moveList, err
}
