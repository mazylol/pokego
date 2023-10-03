package pokemon

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/pokemon"

	_ "github.com/mattn/go-sqlite3"
)

func AddEggGroupToCache(eggGroup pokemon.EggGroup) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&eggGroup)

	_, err = db.Exec("INSERT INTO egg_group (name, data) VALUES (?, ?)", eggGroup.Name, string(jayson))

	return err
}

func GetEggGroupFromCache(name string) (pokemon.EggGroup, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var eggGroup pokemon.EggGroup

	row := db.QueryRow("SELECT data FROM egg_group WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &eggGroup)

	return eggGroup, err
}
