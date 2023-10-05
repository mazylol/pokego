package moves

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/moves"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddCategoryToCache(category moves.Category) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&category)

	_, err = db.Exec("INSERT INTO move_category (name, data) VALUES (?, ?)", category.Name, string(jayson))

	return err
}

func GetCategoryFromCache(name string) (moves.Category, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var category moves.Category

	row := db.QueryRow("SELECT data FROM move_category WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &category)

	return category, err
}
