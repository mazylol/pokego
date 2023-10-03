package pokemon

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mazylol/pokego/types/pokemon"
	"log"
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

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &ability)

	return ability, err
}

func AddAbilityListToCache(abilityList pokemon.AbilityList, count int) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&abilityList)

	_, err = db.Exec("INSERT INTO ability_list (count, data) VALUES (?, ?)", count, string(jayson))

	return err
}

func GetAbilityListFromCache(count int) (pokemon.AbilityList, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var abilityList pokemon.AbilityList

	row := db.QueryRow("SELECT data FROM ability_list WHERE count = ?", count)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &abilityList)

	return abilityList, err
}
