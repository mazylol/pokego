package pokemon

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/pokemon"

	_ "github.com/mattn/go-sqlite3"
)

func AddAbilityToCache(ability pokemon.Ability) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&ability)

	_, err = db.Exec("INSERT INTO ability (name, data) VALUES (?, ?)", ability.Name, string(jayson))

	return err
}

func GetAbilityFromCache(name string) (pokemon.Ability, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var ability pokemon.Ability

	row := db.QueryRow("SELECT data FROM ability WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &ability)

	return ability, err
}
