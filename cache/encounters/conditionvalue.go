package encounters

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/encounters"

	_ "github.com/mattn/go-sqlite3"
)

func AddConditionValueToCache(conditionValue encounters.ConditionValue) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&conditionValue)

	_, err = db.Exec("INSERT INTO encounter_condition_value (name, data) VALUES (?, ?)", conditionValue.Name, string(jayson))

	return err
}

func GetConditionValueFromCache(name string) (encounters.ConditionValue, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var conditionValue encounters.ConditionValue

	row := db.QueryRow("SELECT data FROM encounter_condition_value WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &conditionValue)

	return conditionValue, err
}
