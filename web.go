package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func now() time.Time {
	return time.Now().In(config.Timezone)
}

func beginDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UTC()
}

func endDay(t time.Time) time.Time {
	return beginDay(t).AddDate(0, 0, 1)
}

func HoursHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	q := r.URL.Query()

	var begin, end time.Time
	now := now()

	if len(q["core"]) > 0 && q["core"][0] == "true" {
		begin = time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, now.Location()).UTC()
		end = time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location()).UTC()
	} else {
		begin = beginDay(now)
		end = endDay(now)
	}

	fmt.Printf("Getting hours between %s and %s\n", begin.Format(time.RFC3339), end.Format(time.RFC3339))

	userHours, err := GetHours(begin, end)
	// FIXME: Write http 500?
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"begin": begin.In(config.Timezone),
			"end":   end.In(config.Timezone),
			"users": userHours,
		})
}

func FirstLastsHandler(w http.ResponseWriter, r *http.Request) {
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
