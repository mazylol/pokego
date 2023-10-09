package berries

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/berries"

	_ "github.com/mattn/go-sqlite3"
)

func AddBerryToCache(berry berries.Berry) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&berry)

	_, err = db.Exec("INSERT INTO berry (name, data) VALUES (?, ?)", berry.Name, string(jayson))

	return err
}

func GetBerryFromCache(name string) (berries.Berry, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var berry berries.Berry

	row := db.QueryRow("SELECT data FROM berry WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &berry)

	return berry, err
}
