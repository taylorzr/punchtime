package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HoursHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/usr/local/share/punchtime/hours.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(data))
}

func PunchesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/usr/local/share/punchtime/punches.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(data))
}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/javascript")

	data, err := ioutil.ReadFile("/usr/local/share/punchtime/punchtime.js")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(data))
}
