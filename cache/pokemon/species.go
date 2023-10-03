package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddSpeciesToCache(species pokemon.Species) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&species)

	_, err = db.Exec("INSERT INTO pokemon_species (name, data) VALUES (?, ?)", species.Name, string(jayson))

	return err
}

func GetSpeciesFromCache(name string) (pokemon.Species, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var species pokemon.Species

	row := db.QueryRow("SELECT data FROM pokemon_species WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &species)

	return species, err
}
