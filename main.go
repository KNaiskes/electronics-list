package main

import (
	"net/http"
	"html/template"
	"os"
	"strconv"
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
	http.HandleFunc("/update_component", updateComponentHandler)

	http.Handle("/src/github.com/KNaiskes/electronics-list/static/",
		http.StripPrefix("/src/github.com/KNaiskes/electronics-list/static/",
	http.FileServer(http.Dir("src/github.com/KNaiskes/electronics-list/static/"))))

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
	whichForm := r.FormValue("submit")

	switch(whichForm) {
	case "Submit_Led":
		pieces := r.FormValue("pieces")
		color  := r.FormValue("color")

		piecesInt, _ := strconv.Atoi(pieces)

		led := database.Leds{Piece: piecesInt, Color: color}
		database.NewComponentDB(led)
		break
	case "Submit_Board":
		pieces   := r.FormValue("pieces")
		name     := r.FormValue("name")
		ethernet := r.FormValue("ethernet")
		wifi     := r.FormValue("wifi")
		version  := r.FormValue("version")

		piecesInt, _    := strconv.Atoi(pieces)
		ethernetBool, _ := strconv.ParseBool(ethernet)
		wifiBool, _     := strconv.ParseBool(wifi)

		board := database.Board{Piece: piecesInt, Name: name,
				HasEthernet: ethernetBool, HasWifi: wifiBool,
				Version: version}
		database.NewComponentDB(board)
		break
	case "Submit_Jumper":
		pieces := r.FormValue("pieces")
		cm     := r.FormValue("cm")
		jtype  := r.FormValue("type")

		piecesInt, _ := strconv.Atoi(pieces)
		cmFloat, _   := strconv.ParseFloat(cm, 32)

		jumperWire := database.JumperWire{Piece: piecesInt,
				Cm: cmFloat, Jtype: jtype}
		database.NewComponentDB(jumperWire)
		break
	case "Submit_resistor":
		pieces := r.FormValue("pieces")
		value  := r.FormValue("value")

		piecesInt, _  := strconv.Atoi(pieces)
		valueFloat, _ := strconv.ParseFloat(value, 32)

		resistor := database.Resistor{Piece: piecesInt,
				Value: valueFloat}

		database.NewComponentDB(resistor)
		break
	}

	tmpl.ExecuteTemplate(w, "add_component.html", nil)
}

func removeComponentHandler(w http.ResponseWriter, r *http.Request) {
	whichForm := r.FormValue("submit")

	switch(whichForm) {
		case "Submit_Led":
			color := r.FormValue("color")
			led := database.Leds{Color: color}
			database.RemoveComponentDB(led)
			break
		case "Submit_Board":
			name := r.FormValue("name")
			board := database.Board{Name: name}
			database.RemoveComponentDB(board)
			break
		case "Submit_Jumper":
			jtype := r.FormValue("type")
			jumperWire := database.JumperWire{Jtype: jtype}
			database.RemoveComponentDB(jumperWire)
			break
		case "Submit_resistor":
			value := r.FormValue("value")
			valueFloat, _ := strconv.ParseFloat(value, 32)
			resistor := database.Resistor{Value: valueFloat}
			database.RemoveComponentDB(resistor)
			break
		}

	tmpl.ExecuteTemplate(w, "remove_component.html", nil)
}

func updateComponentHandler(w http.ResponseWriter, r *http.Request) {
	whichForm := r.FormValue("submit")

	switch(whichForm) {
	case "Submit_Led":
		color       := r.FormValue("color")
		updateColor := r.FormValue("updatedColor")
		pieces      := r.FormValue("pieces")

		piecesInt, _ := strconv.Atoi(pieces)

		led := database.Leds{Piece: piecesInt, Color: updateColor}
		database.UpdateComponent(led, color)
		break
	case "Submit_Board":
		name       := r.FormValue("name")
		pieces     := r.FormValue("pieces")
		ethernet   := r.FormValue("ethernet")
		wifi       := r.FormValue("wifi")
		version    := r.FormValue("version")
		updateName := r.FormValue("updatedName")

		piecesInt, _    := strconv.Atoi(pieces)
		ethernetBool, _ := strconv.ParseBool(ethernet)
		wifiBool, _     := strconv.ParseBool(wifi)

		board := database.Board{Piece: piecesInt, Name: updateName,
				HasEthernet: ethernetBool, HasWifi: wifiBool,
				Version: version}
		database.UpdateComponent(board, name)
		break
	case "Submit_Jumper":
		updatedType := r.FormValue("updateType")
		pieces      := r.FormValue("pieces")
		cm          := r.FormValue("cm")
		jtype       := r.FormValue("type")

		piecesInt, _ := strconv.Atoi(pieces)
		cmFloat, _   := strconv.ParseFloat(cm, 32)

		jumperWire := database.JumperWire{Piece: piecesInt,
				Cm: cmFloat, Jtype: updatedType}
		database.UpdateComponent(jumperWire, jtype)
		break
	case "Submit_resistor":
		pieces := r.FormValue("pieces")
		value  := r.FormValue("value")
		updatedValue := r.FormValue("updateValue")

		piecesInt, _ := strconv.Atoi(pieces)
		updatedValueFloat, _ := strconv.ParseFloat(updatedValue, 32)

		//pieces := r.FormValue("pieces")
		//value  := r.FormValue("value")
		//updatedValue := r.FormValue("updateValue")

		//piecesInt, _          := strconv.Atoi(pieces)
		////valueFloat, _         := strconv.ParseFloat(value, 32)
		//updadateValueFloat, _ := strconv.ParseFloat(updatedValue, 32)

		resistor := database.Resistor{Piece: piecesInt,
				Value: updatedValueFloat}

		database.UpdateComponent(resistor, value)
		break
	}

	tmpl.ExecuteTemplate(w, "modify_component.html", nil)
}
