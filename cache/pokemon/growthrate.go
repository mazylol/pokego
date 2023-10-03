package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddGrowthRateToCache(growthRate pokemon.GrowthRate) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&growthRate)

	_, err = db.Exec("INSERT INTO growth_rate (name, data) VALUES (?, ?)", growthRate.Name, string(jayson))

	return err
}

func GetGrowthRateFromCache(name string) (pokemon.GrowthRate, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var growthRate pokemon.GrowthRate

	row := db.QueryRow("SELECT data FROM growth_rate WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &growthRate)

	return growthRate, err
}
