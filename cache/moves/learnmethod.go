package moves

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/moves"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddLearnMethodToCache(learnMethod moves.LearnMethod) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&learnMethod)

	_, err = db.Exec("INSERT INTO move_learn_method (name, data) VALUES (?, ?)", learnMethod.Name, string(jayson))

	return err
}

func GetLearnMethodFromCache(name string) (moves.LearnMethod, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var learnMethod moves.LearnMethod

	row := db.QueryRow("SELECT data FROM move_learn_method WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &learnMethod)

	return learnMethod, err
}
