package main

import (
	"net/http"
	"html/template"
	"os"
	"github.com/KNaiskes/electronics-list/database"
)

type Components struct {
	// any new component type must be added here
	Leds        interface{}
	Board       interface{}
	Jumperwire  interface{}
	Resistor    interface{}
}

var htmlDir = "src/github.com/KNaiskes/electronics-list/static/html/index.html"
var temp = "src/github.com/KNaiskes/electronics-list/static/html/components.html"

func main() {

	if _, err := os.Stat(database.Dbdir); os.IsNotExist(err) {
		os.MkdirAll(database.Dbdir, 0700)
		database.CreateDB()
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/components", componentsHandler)

	http.Handle("/src/github.com/KNaiskes/electronics-list/static/css/",
		http.StripPrefix("/src/github.com/KNaiskes/electronics-list/static/css/",
	http.FileServer(http.Dir("src/github.com/KNaiskes/electronics-list/static/css/"))))

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(htmlDir))
	tmpl.Execute(w, nil)
}

func componentsHandler(w http.ResponseWriter, r *http.Request) {
	var(
		l = database.Leds{}
		b = database.Board{}
		j = database.JumperWire{}
		re = database.Resistor{}
	)

	components := Components{ Leds: database.ListComponent(l),
		Board: database.ListComponent(b),
		Jumperwire: database.ListComponent(j),
		Resistor: database.ListComponent(re) }

	tmpl := template.Must(template.ParseFiles(temp))
	tmpl.Execute(w, components)
}
