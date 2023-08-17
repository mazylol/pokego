package pokego

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mazylol/pokego/structs"
	"log"
	"os"
)

func init() {
	if _, err := os.Stat("./pokego.db"); err == nil {
		return
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Creating database...")
		create()
	} else {
		log.Fatal(err)
	}
}

func create() {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	sqlStmt := "CREATE TABLE IF NOT EXISTS pokemon (name VARCHAR(50) PRIMARY KEY, data TEXT[] NOT NULL)"

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func addPokemonToCache(pokemon structs.Pokemon) error {
	//start()

	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&pokemon)

	_, err = db.Exec("INSERT INTO pokemon (name, data) VALUES (?, ?)", pokemon.Name, string(jayson))

	return err
}

func getPokemonFromCache(name string) (structs.Pokemon, error) {
	//start()

	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var pokemon structs.Pokemon

	row := db.QueryRow("SELECT data FROM pokemon WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &pokemon)

	return pokemon, err
}
