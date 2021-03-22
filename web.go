package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed "hours.html"
var hours_html string

//go:embed "punches.html"
var punches_html string

//go:embed "punchtime.js"
var punchtime_js string

func HoursHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hours_html)
}

func PunchesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, punches_html)
}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/javascript")

	fmt.Fprint(w, punchtime_js)
}
