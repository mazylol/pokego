package moves

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/moves"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddDamageClassToCache(damageClass moves.DamageClass) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&damageClass)

	_, err = db.Exec("INSERT INTO move_damage_class (name, data) VALUES (?, ?)", damageClass.Name, string(jayson))

	return err
}

func GetDamageClassFromCache(name string) (moves.DamageClass, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var damageClass moves.DamageClass

	row := db.QueryRow("SELECT data FROM move_damage_class WHERE name = ?", name)

	var data []byte
	err = row.Scan(&data)

	err = json.Unmarshal(data, &damageClass)

	return damageClass, err
}
