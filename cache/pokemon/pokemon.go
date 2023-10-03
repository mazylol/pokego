package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddPokemonToCache(pokemon pokemon.Pokemon) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&pokemon)

	_, err = db.Exec("INSERT INTO pokemon (name, data) VALUES (?, ?)", pokemon.Name, string(jayson))

	return err
}

func GetPokemonFromCache(name string) (pokemon.Pokemon, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var pokemon pokemon.Pokemon

	row := db.QueryRow("SELECT data FROM pokemon WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &pokemon)

	return pokemon, err
}
