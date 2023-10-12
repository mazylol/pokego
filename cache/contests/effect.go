package contests

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/contests"

	_ "github.com/mattn/go-sqlite3"
)

func AddEffectToCache(effect contests.Effect) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&effect)

	_, err = db.Exec("INSERT INTO contest_effect (id, data) VALUES (?, ?)", effect.Id, string(jayson))

	return err
}

func GetEffectFromCache(id int) (contests.Effect, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var effect contests.Effect

	row := db.QueryRow("SELECT data FROM contest_effect WHERE id = ?", id)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &effect)

	return effect, err
}
