package moves

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/moves"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddTargetToCache(target moves.Target) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&target)

	_, err = db.Exec("INSERT INTO move_target (name, data) VALUES (?, ?)", target.Name, string(jayson))

	return err
}

func GetTargetFromCache(name string) (moves.Target, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var target moves.Target

	row := db.QueryRow("SELECT data FROM move_target WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &target)

	return target, err
}
