package pokemon

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/pokemon"

	_ "github.com/mattn/go-sqlite3"
)

func AddCharacteristicToCache(characteristic pokemon.Characteristic) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&characteristic)

	_, err = db.Exec("INSERT INTO characteristic (id, data) VALUES (?, ?)", characteristic.Id, string(jayson))

	return err
}

func GetCharacteristicFromCache(id int) (pokemon.Characteristic, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var characteristic pokemon.Characteristic

	row := db.QueryRow("SELECT data FROM characteristic WHERE id = ?", id)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &characteristic)

	return characteristic, err
}
