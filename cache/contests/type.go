package contests

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/contests"

	_ "github.com/mattn/go-sqlite3"
)

func AddTypeToCache(tipe contests.Type) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&tipe)

	_, err = db.Exec("INSERT INTO contest_type (name, data) VALUES (?, ?)", tipe.Name, string(jayson))

	return err
}

func GetTypeFromCache(name string) (contests.Type, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var tipe contests.Type

	row := db.QueryRow("SELECT data FROM contest_type WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &tipe)

	return tipe, err
}
