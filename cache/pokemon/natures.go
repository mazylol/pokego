package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddNatureToCache(nature pokemon.Nature) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&nature)

	_, err = db.Exec("INSERT INTO nature (name, data) VALUES (?, ?)", nature.Name, string(jayson))

	return err
}

func GetNatureFromCache(name string) (pokemon.Nature, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var nature pokemon.Nature

	row := db.QueryRow("SELECT data FROM nature WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &nature)

	return nature, err
}
