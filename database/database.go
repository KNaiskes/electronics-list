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
	statement.Exec(l.Color)

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
	statement.Exec(l.Piece, l.Color, "blue")

	if err != nil {
		log.Fatal(err)
	}
}

func (b Board) AddComponent() {
	const query = `INSERT INTO boards(piece, name, ethernet, wifi, version) VALUES(?, ?, ?, ?, ?)`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(b.Piece, b.Name, b.HasEthernet, b.HasWifi, b.Version)

	if err != nil {
		log.Fatal(err)
	}
}

func (b Board) DeleteComponent() {
	const query = `DELETE FROM boards WHERE name = ?`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(b.Name)

	if err != nil {
		log.Fatal(err)
	}
}

func (b Board) ModifyComponent() {
	const query = `UPDATE boards SET piece = ?, name = ?, ethernet = ?, wifi = ?, version = ? WHERE name = ?`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(b.Piece, b.Name, b.HasEthernet, b.HasWifi, b.Version, "Raspberry Pi") // just for testing

	if err != nil {
		log.Fatal(err)
	}
}

func (j JumperWire) AddComponent() {
	const query = `INSERT INTO jumberwires(piece, cm, type) VALUES(?, ?, ?)`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(j.Piece, j.Cm, j.Jtype)

	if err != nil {
		log.Fatal(err)
	}
}

func (j JumperWire) DeleteComponent() {
	const query = `DELETE FROM jumberwires WHERE type = ?`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(j.Jtype)

	if err != nil {
		log.Fatal(err)
	}
}

func (j JumperWire) ModifyComponent() {
	const query = `UPDATE jumberwires SET piece = ?, cm = ?, type = ?`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(j.Piece, j.Cm, j.Jtype)

	if err != nil {
		log.Fatal(err)
	}
}

func (r Resistor) AddComponent() {
	const query = `INSERT INTO resistors(piece, value) VALUES(?,?)`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(r.Piece, r.Value)

	if err != nil {
		log.Fatal(err)
	}
}

func (r Resistor) DeleteComponent() {
	const query = `DELETE FROM resistors WHERE value = ?`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(r.Value)

	if err != nil {
		log.Fatal(err)
	}
}

func (r Resistor) ModifyComponent() {
	const query = `UPDATE resistors SET piece = ?, value = ?`

	db, err := sql.Open(driverDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(query)
	statement.Exec(r.Piece, r.Value)

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
