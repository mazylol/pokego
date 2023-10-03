package pokemon

import (
	"database/sql"
	"encoding/json"
	"github.com/mazylol/pokego/types/pokemon"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddPokeathlonStatToCache(pokeathlonStat pokemon.PokeathlonStat) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&pokeathlonStat)

	_, err = db.Exec("INSERT INTO pokeathlon_stat (name, data) VALUES (?, ?)", pokeathlonStat.Name, string(jayson))

	return err
}

func GetPokeathlonStatFromCache(name string) (pokemon.PokeathlonStat, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var pokeathlonStat pokemon.PokeathlonStat

	row := db.QueryRow("SELECT data FROM pokeathlon_stat WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &pokeathlonStat)

	return pokeathlonStat, err
}
