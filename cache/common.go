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
		CREATE TABLE IF NOT EXISTS berry_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS berry_firmness_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS berry_flavor_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);

		CREATE TABLE IF NOT EXISTS contest_type_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS contest_effect_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS super_contest_effect_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);

		CREATE TABLE IF NOT EXISTS encounter_method_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS encounter_condition_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS encounter_condition_value_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);

		CREATE TABLE IF NOT EXISTS evolution_chain_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS evolution_trigger_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);

		CREATE TABLE IF NOT EXISTS generation_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS pokedex_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS version_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS version_group_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);

		CREATE TABLE IF NOT EXISTS item_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS item_attribute_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS item_category_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS item_fling_effect_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS item_pocket_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);

		CREATE TABLE IF NOT EXISTS location_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS location_area_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS pal_park_area_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS region_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);

		CREATE TABLE IF NOT EXISTS machine_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);

		CREATE TABLE IF NOT EXISTS move_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS move_ailment_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS move_battle_style_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS move_category_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS move_damage_class_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS move_learn_method_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS move_target_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);

	    CREATE TABLE IF NOT EXISTS ability_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS characteristic_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS egg_group_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS gender_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS growth_rate_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS nature_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS pokeathlon_stat_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS pokemon_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS pokemon_color_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS pokemon_form_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS pokemon_habitat_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS pokemon_shape_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS pokemon_species_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS stat_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS type_list (count INTEGER PRIMARY KEY, data TEXT[] NOT NULL);

		CREATE TABLE IF NOT EXISTS pokemon (name VARCHAR(50) PRIMARY KEY, data TEXT[] NOT NULL);
		CREATE TABLE IF NOT EXISTS ability (name VARCHAR(50) PRIMARY KEY, data TEXT[] NOT NULL);

		CREATE TABLE IF NOT EXISTS move (name VARCHAR(50) PRIMARY KEY, data TEXT[] NOT NULL);

`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}
