package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const driver = "sqlite3"
const dbDir  = "src/github.com/KNaiskes/electronics-list/db/components.db"

func newTable(driver string, query string) {
	db, err := sql.Open(driver, dbDir)

	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

}
