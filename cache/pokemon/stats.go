package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddStatToCache(stat pokemon.Stat) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&stat)

	_, err = db.Exec("INSERT INTO stat (name, data) VALUES (?, ?)", stat.Name, string(jayson))

	return err
}

func GetStatFromCache(name string) (pokemon.Stat, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var stat pokemon.Stat

	row := db.QueryRow("SELECT data FROM stat WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &stat)

	return stat, err
}
