package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const driverDB = "sqlite3"
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

func CreateDB() {
	const queryLed = `CREATE TABLE IF NOT EXISTS leds(id INTEGER PRIMARY KEY,
		  piece INTEGER, color TEXT)`
	const queryResistor = `CREATE TABLE IF NOT EXISTS resistors
			       (id INTEGER PRIMARY KEY, piece INTEGER, value REAL)`
	const queryJumperWire = `CREATE TABLE IF NOT EXISTS jumberwires
				 (id INTEGER PRIMARY KEY, piece INTEGER, cm REAL,
				 type TEXT)`
	const queryBoard = `CREATE TABLE IF NOT EXISTS boards
			    (id INTEGER PRIMARY KEY, piece INTEGER, name TEXT,
			    ethernet TEXT, wifi TEXT, version TEXT)`

	newTable(driverDB, queryLed)
	newTable(driverDB, queryResistor)
	newTable(driverDB, queryJumperWire)
	newTable(driverDB, queryBoard)
}
