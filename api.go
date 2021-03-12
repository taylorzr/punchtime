package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ApiPunchesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	username := strings.TrimPrefix(r.URL.Path, "/api/punches/")

	punches := PunchesFor(username)

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(punches)
}

func ApiHoursHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	q := r.URL.Query()

	day := 0
	if len(q["day"]) > 0 {
		var err error
		day, err = strconv.Atoi(q["day"][0])
		if err != nil {
			log.Fatal(err)
		}
		if day > 0 {
			log.Fatal("Day must be less than or equal to 0")
		}
	}

	now := now().AddDate(0, 0, day)
	dayLog := GetDayLog(now)

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(dayLog)
}

func ApiFirstLastsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	now := now()
	begin := beginDay(now)
	end := endDay(now)

	firstLasts, err := GetFirstLasts(begin, end)
	// FIXME: Write http 500?
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"begin":       begin.In(config.Timezone),
			"end":         end.In(config.Timezone),
			"first_lasts": firstLasts,
		})
}
