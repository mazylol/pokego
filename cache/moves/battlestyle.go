package moves

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/mazylol/pokego/types/moves"

	_ "github.com/mattn/go-sqlite3"
)

func AddBattleStyleToCache(battleStyle moves.BattleStyle) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&battleStyle)

	_, err = db.Exec("INSERT INTO move_battle_style (name, data) VALUES (?, ?)", battleStyle.Name, string(jayson))

	return err
}

func GetBattleStyleFromCache(name string) (moves.BattleStyle, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var battleStyle moves.BattleStyle

	row := db.QueryRow("SELECT data FROM move_battle_style WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &battleStyle)

	return battleStyle, err
}
