package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"
)

func AddFormToCache(form pokemon.Form) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&form)

	_, err = db.Exec("INSERT INTO pokemon_form (name, data) VALUES (?, ?)", form.Name, string(jayson))

	return err
}

func GetFormFromCache(name string) (pokemon.Form, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var form pokemon.Form

	row := db.QueryRow("SELECT data FROM pokemon_form WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &form)

	return form, err
}
