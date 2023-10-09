package berries

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/berries"

	_ "github.com/mattn/go-sqlite3"
)

func AddFirmnessToCache(firmness berries.Firmness) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&firmness)

	_, err = db.Exec("INSERT INTO berry_firmness (name, data) VALUES (?, ?)", firmness.Name, string(jayson))

	return err
}

func GetFirmnessFromCache(name string) (berries.Firmness, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var firmness berries.Firmness

	row := db.QueryRow("SELECT data FROM berry_firmness WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &firmness)

	return firmness, err
}
