package encounters

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/encounters"

	_ "github.com/mattn/go-sqlite3"
)

func AddConditionToCache(condition encounters.Condition) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&condition)

	_, err = db.Exec("INSERT INTO encounter_condition (name, data) VALUES (?, ?)", condition.Name, string(jayson))

	return err
}

func GetConditionFromCache(name string) (encounters.Condition, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var condition encounters.Condition

	row := db.QueryRow("SELECT data FROM encounter_condition WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &condition)

	return condition, err
}
