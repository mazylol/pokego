package cache

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

func Initialize() {
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
