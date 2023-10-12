package contests

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/contests"

	_ "github.com/mattn/go-sqlite3"
)

func AddSuperEffectToCache(superEffect contests.SuperEffect) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&superEffect)

	_, err = db.Exec("INSERT INTO super_contest_effect (id, data) VALUES (?, ?)", superEffect.Id, string(jayson))

	return err
}

func GetSuperEffectFromCache(id int) (contests.SuperEffect, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var superEffect contests.SuperEffect

	row := db.QueryRow("SELECT data FROM super_contest_effect WHERE id = ?", id)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &superEffect)

	return superEffect, err
}
