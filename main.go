package main

import (
	"net/http"
	"html/template"
	"os"
	"github.com/KNaiskes/electronics-list/database"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("src/github.com/KNaiskes/electronics-list/static/html/*.html"))
}

type Components struct {
	// any new component type must be added here
	Leds        interface{}
	Board       interface{}
	Jumperwire  interface{}
	Resistor    interface{}
}

func main() {

	if _, err := os.Stat(database.Dbdir); os.IsNotExist(err) {
		os.MkdirAll(database.Dbdir, 0700)
		database.CreateDB()
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/components", componentsHandler)
	http.HandleFunc("/add_component", addComponentHandler)
	http.HandleFunc("/remove_component", removeComponentHandler)

	http.Handle("/src/github.com/KNaiskes/electronics-list/static/css/",
		http.StripPrefix("/src/github.com/KNaiskes/electronics-list/static/css/",
	http.FileServer(http.Dir("src/github.com/KNaiskes/electronics-list/static/css/"))))

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "index.html", nil)
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

		tmpl.ExecuteTemplate(w, "components.html", components)
}

func addComponentHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "add_component.html", nil)
}

func removeComponentHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "remove_component.html", nil)
}
