package berries

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/berries"

	_ "github.com/mattn/go-sqlite3"
)

func AddFlavorToCache(flavor berries.Flavor) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&flavor)

	_, err = db.Exec("INSERT INTO berry_flavor (name, data) VALUES (?, ?)", flavor.Name, string(jayson))

	return err
}

func GetFlavorFromCache(name string) (berries.Flavor, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var flavor berries.Flavor

	row := db.QueryRow("SELECT data FROM berry_flavor WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &flavor)

	return flavor, err
}
