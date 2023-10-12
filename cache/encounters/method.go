package encounters

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/encounters"

	_ "github.com/mattn/go-sqlite3"
)

func AddMethodToCache(method encounters.Method) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&method)

	_, err = db.Exec("INSERT INTO encounter_method (name, data) VALUES (?, ?)", method.Name, string(jayson))

	return err
}

func GetMethodFromCache(name string) (encounters.Method, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var method encounters.Method

	row := db.QueryRow("SELECT data FROM encounter_method WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &method)

	return method, err
}
