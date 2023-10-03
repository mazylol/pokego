package pokemon

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/pokemon"

	_ "github.com/mattn/go-sqlite3"
)

func AddShapeToCache(shape pokemon.Shape) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&shape)

	_, err = db.Exec("INSERT INTO pokemon_shape (name, data) VALUES (?, ?)", shape.Name, string(jayson))

	return err
}

func GetShapeFromCache(name string) (pokemon.Shape, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var shape pokemon.Shape

	row := db.QueryRow("SELECT data FROM pokemon_shape WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &shape)

	return shape, err
}
