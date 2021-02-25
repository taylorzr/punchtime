package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func HoursHandler(w http.ResponseWriter, r *http.Request) {
	chicago, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Fatal(err)
	}

	chicagoNow := time.Now().In(chicago).AddDate(0, 0, -3)
	chicagoBegin := time.Date(chicagoNow.Year(), chicagoNow.Month(), chicagoNow.Day(), 0, 0, 0, 0, chicagoNow.Location())
	chicagoEnd := chicagoBegin.AddDate(0, 0, 1)
	begin := chicagoBegin.UTC()
	end := chicagoEnd.UTC()

	fmt.Printf("%s -> %s\n", chicagoBegin.Format(time.RFC3339), chicagoEnd.Format(time.RFC3339))
	fmt.Printf("%s -> %s\n", begin.Format(time.RFC3339), end.Format(time.RFC3339))

	userHours, err := GetHours(begin, end)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"time":  time.Now(),
			"users": userHours,
		})
}
