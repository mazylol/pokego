package pokego

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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

	sqlStmt := `
		CREATE TABLE IF NOT EXISTS pokemon (name VARCHAR(50) PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS pokemon_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS move (name VARCHAR(50) PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS move_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL)
`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func addPokemonToCache(pokemon Pokemon) error {
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

func getPokemonFromCache(name string) (Pokemon, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var pokemon Pokemon

	row := db.QueryRow("SELECT data FROM pokemon WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &pokemon)

	return pokemon, err
}

func addPokemonListToCache(pokemonList PokemonList, count int) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&pokemonList)

	_, err = db.Exec("INSERT INTO pokemon_list (count, data) VALUES (?, ?)", count, string(jayson))

	return err
}

func getPokemonListFromCache(count int) (PokemonList, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var pokemonList PokemonList

	row := db.QueryRow("SELECT data FROM pokemon_list WHERE count = ?", count)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &pokemonList)

	return pokemonList, err
}

func addMoveToCache(move Move) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&move)

	_, err = db.Exec("INSERT INTO move (name, data) VALUES (?, ?)", move.Name, string(jayson))

	return err
}

func getMoveFromCache(name string) (Move, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var move Move

	row := db.QueryRow("SELECT data FROM move WHERE name = ?", name)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &move)

	return move, err
}

func addMoveListToCache(moveList MoveList, count int) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&moveList)

	_, err = db.Exec("INSERT INTO move_list (count, data) VALUES (?, ?)", count, string(jayson))

	return err
}

func getMoveListFromCache(count int) (MoveList, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var moveList MoveList

	row := db.QueryRow("SELECT data FROM move_list WHERE count = ?", count)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &moveList)

	return moveList, err
}
