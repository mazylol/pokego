package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddColorToCache(color pokemon.Color) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&color)

	_, err = db.Exec("INSERT INTO pokemon_color (name, data) VALUES (?, ?)", color.Name, string(jayson))

	return err
}

func GetColorFromCache(name string) (pokemon.Color, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var color pokemon.Color

	row := db.QueryRow("SELECT data FROM pokemon_color WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &color)

	return color, err
}
