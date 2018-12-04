package main

import (
	"net/http"
	"html/template"
)

type Resistor struct {
	piece int
	value float32
}

type Leds struct {
	piece int
	color string
}

type boards struct {
	piece	    int
	name	    string
	hasEthernet bool
	hasWifi	    bool
	version     string
}

type jumperWires struct {
	piece int
	cm    float32
	jtype string
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}
