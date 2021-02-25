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

var db *sqlx.DB
var slackClient *slack.Client

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "serve",
				Action: func(c *cli.Context) error {
					var err error
					db, err = sqlx.Connect("sqlite", "/usr/local/share/punchtime/punchtime.db")
					if err != nil {
						return err
					}

					http.HandleFunc("/hours", HoursHandler)
					fmt.Println("Server started at port 8080")
					return http.ListenAndServe("127.0.0.1:8080", nil)
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

					slackClient = slack.New(os.Getenv("SLACK_TOKEN"))

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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
