package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"

	_ "modernc.org/sqlite"
)

var config Punchtime

type Punchtime struct {
	Timezone *time.Location
	DB       *sqlx.DB
	Slack    *slack.Client
}

func main() {
	db, err := sqlx.Connect("sqlite", "/usr/local/share/punchtime/punchtime.db")
	if err != nil {
		log.Fatal(err)
	}

	chicago, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Fatal(err)
	}

	slackClient := slack.New(os.Getenv("SLACK_TOKEN"))

	config = Punchtime{
		Timezone: chicago,
		DB:       db,
		Slack:    slackClient,
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "serve",
				Action: func(c *cli.Context) error {
					http.HandleFunc("/punchtime.js", JsHandler)
					http.HandleFunc("/hours", HoursHandler)
					http.HandleFunc("/punches/", PunchesHandler)
					http.HandleFunc("/api/hours", ApiHoursHandler)
					http.HandleFunc("/api/firstlasts", ApiFirstLastsHandler)
					http.HandleFunc("/api/punches/", config.ApiPunchesHandler)
					fmt.Println("Server started at port 8081")
					location := os.Getenv("LOCATION")
					if location == "" {
						location = ":8081"
					}
					return http.ListenAndServe(location, nil)
				},
			},
			{
				Name: "iterate",
				Action: func(c *cli.Context) error {
					var err error
					db, err = sqlx.Connect("sqlite", "/usr/local/share/punchtime/punchtime.db")
					if err != nil {
						return err
					}

					users, err := GetUsers()
					if err != nil {
						return err
					}

					for _, user := range users {
						err := user.MaybePunch(time.Now().UTC())
						if err != nil {
							return err
						}
					}

					return nil
				},
			},
			{
				Name: "test",
				Action: func(c *cli.Context) error {
					today := time.Now()
					tomorrow := today.AddDate(0, 0, 1)
					fmt.Println(today)
					fmt.Println(today.UTC())
					fmt.Println(time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location()))
					fmt.Println(time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location()).UTC())
					fmt.Println(time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, today.Location()).UTC())
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
