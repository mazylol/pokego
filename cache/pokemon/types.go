package pokemon

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/pokemon"

	_ "github.com/mattn/go-sqlite3"
)

func AddTypeToCache(tipe pokemon.Type) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&tipe)

	_, err = db.Exec("INSERT INTO type (name, data) VALUES (?, ?)", tipe.Name, string(jayson))

	return err
}

func GetTypeFromCache(name string) (pokemon.Type, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var tipe pokemon.Type

	row := db.QueryRow("SELECT data FROM type WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &tipe)

	return tipe, err
}
