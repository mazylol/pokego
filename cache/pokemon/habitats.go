package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddHabitatToCache(habitat pokemon.Habitat) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&habitat)

	_, err = db.Exec("INSERT INTO pokemon_habitat (name, data) VALUES (?, ?)", habitat.Name, string(jayson))

	return err
}

func GetHabitatFromCache(name string) (pokemon.Habitat, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var habitat pokemon.Habitat

	row := db.QueryRow("SELECT data FROM pokemon_habitat WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &habitat)

	return habitat, err
}
