package moves

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/moves"

	_ "github.com/mattn/go-sqlite3"
)

func AddAilmentToCache(ailment moves.Ailment) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&ailment)

	_, err = db.Exec("INSERT INTO move_ailment (name, data) VALUES (?, ?)", ailment.Name, string(jayson))

	return err
}

func GetAilmentFromCache(name string) (moves.Ailment, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var ailment moves.Ailment

	row := db.QueryRow("SELECT data FROM move_ailment WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &ailment)

	return ailment, err
}
