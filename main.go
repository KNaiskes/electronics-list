package main

import (
	"net/http"
	"html/template"
	"os"
	"github.com/KNaiskes/electronics-list/resistor"
	//"github.com/KNaiskes/electronics-list/led"
	//"github.com/KNaiskes/electronics-list/board"
	//"github.com/KNaiskes/electronics-list/jumberWire"
	"github.com/KNaiskes/electronics-list/database"
)

var htmlDir = "src/github.com/KNaiskes/electronics-list/static/html/index.html"

func main() {

	if _, err := os.Stat(database.Dbdir); os.IsNotExist(err) {
		os.MkdirAll(database.Dbdir, 0700)
		database.CreateDB()
	}

	http.HandleFunc("/", indexHandler)

	http.Handle("/src/github.com/KNaiskes/electronics-list/static/css/",
		http.StripPrefix("/src/github.com/KNaiskes/electronics-list/static/css/",
	http.FileServer(http.Dir("src/github.com/KNaiskes/electronics-list/static/css/"))))

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	kr := resistor.Addresistor(32, 56.44)
	tmpl := template.Must(template.ParseFiles(htmlDir))
	tmpl.Execute(w, kr)
}
