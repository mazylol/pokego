package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddLocationAreaToCache(locationArea []pokemon.LocationArea, name string) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&locationArea)

	_, err = db.Exec("INSERT INTO pokemon_location_area (name, data) VALUES (?, ?)", name, string(jayson))

	return err
}

func GetLocationAreaFromCache(name string) ([]pokemon.LocationArea, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var locationArea []pokemon.LocationArea

	row := db.QueryRow("SELECT data FROM pokemon_location_area WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &locationArea)

	return locationArea, err
}
