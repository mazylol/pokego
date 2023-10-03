package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddGenderToCache(gender pokemon.Gender) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&gender)

	_, err = db.Exec("INSERT INTO gender (name, data) VALUES (?, ?)", gender.Name, string(jayson))

	return err
}

func GetGenderFromCache(name string) (pokemon.Gender, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var gender pokemon.Gender

	row := db.QueryRow("SELECT data FROM gender WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &gender)

	return gender, err
}
