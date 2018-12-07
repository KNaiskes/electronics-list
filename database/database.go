package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Leds struct {
	Piece int
	Color string
}

type Board struct {
	Piece	    int
	Name	    string
	HasEthernet bool
	HasWifi	    bool
	Version     string
}

type JumperWire struct {
	Piece int
	Cm    float32
	Jtype string
}

type Resistor struct {
	Piece int
	Value float32
}

func Addresistor(p int, v float32) Resistor {
	return Resistor{Piece: p, Value: v}
}

type DatabaseInterface interface {
	AddComponent()
	DeleteComponent()
	ModifyComponent()
}

const driverDB = "sqlite3"
const Dbdir  = "src/github.com/KNaiskes/electronics-list/db/"

const dbName = Dbdir + "components.db"

func newTable(driver string, query string) {
	db, err := sql.Open(driver, dbName)

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

func (l Leds) AddComponent() {
	const query = `INSERT INTO leds(piece, color) VALUES(?,?)`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(l.Piece, l.Color)

	if err != nil {
		log.Fatal(err)
	}
}

func (l Leds) DeleteComponent() {
	const query = `DELETE FROM leds WHERE COLOR = ?`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(l.Color) // delete by color for now just for testing

	if err != nil {
		log.Fatal(err)
	}
}

func (l Leds) ModifyComponent() {
	const query = `UPDATE leds SET piece = ?, color = ? WHERE color = ?`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(l.Piece, l.Color, "blue") // modify by color for now just for testing

	if err != nil {
		log.Fatal(err)
	}
}

// implementation of DatabaseInterface interface

func NewComponentDB(d DatabaseInterface) {
	d.AddComponent()
}

func RemoveComponentDB(d DatabaseInterface) {
	d.DeleteComponent()
}

func UpdateComponent(d DatabaseInterface) {
	d.ModifyComponent()
}
