package resource

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/mazylol/pokego/types/resource"
	"log"
	"strings"
)

func getTableName(endpoint string) string {
	return strings.ReplaceAll(fmt.Sprintf("%s_list", endpoint), "-", "_")
}

func AddResourceListToCache(resourceList resource.List, endpoint string, count int) error {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	jayson, err := json.Marshal(&resourceList)

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (count, data) VALUES (?, ?)", getTableName(endpoint)), count, string(jayson))

	return err
}

func GetResourceListFromCache(endpoint string, count int) (resource.List, error) {
	db, err := sql.Open("sqlite3", "./pokego.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)

	var resourceList resource.List

	row := db.QueryRow("SELECT data FROM ? WHERE count = ?", getTableName(endpoint), count)

	var data string
	err = row.Scan(&data)

	err = json.Unmarshal([]byte(data), &resourceList)

	return resourceList, err
}
