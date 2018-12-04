package main

import (
	"net/http"
	"html/template"
)

type Resistor struct {
	Piece int
	Value float32
}

type Leds struct {
	Piece int
	Color string
}

type Boards struct {
	Piece	    int
	Name	    string
	HasEthernet bool
	HasWifi	    bool
	Version     string
}

type JumperWires struct {
	Piece int
	Cm    float32
	Jtype string
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}
