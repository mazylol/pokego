package moves

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/moves"

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

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &move)

	return move, err
}
