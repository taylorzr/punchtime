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

func begin(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UTC()
}

func end(t time.Time) time.Time {
	return begin(t).AddDate(0, 0, 1)
}

func HoursHandler(w http.ResponseWriter, r *http.Request) {
	now := now()
	begin := begin(now)
	end := end(now)

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
	now := now()
	begin := begin(now)
	end := end(now)

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
