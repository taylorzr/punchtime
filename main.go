package main

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/slack-go/slack"

	// _ "github.com/mattn/go-sqlite3" Some problem with CGO on raspberry pi
	_ "modernc.org/sqlite"
)

var db *sqlx.DB
var slackClient *slack.Client

func main() {
	var err error
	db, err = sqlx.Connect("sqlite", "/usr/local/share/punchtime/punchtime.db")
	if err != nil {
		log.Fatal(err)
	}

	slackClient = slack.New(os.Getenv("SLACK_TOKEN"))

	users, err := GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	chicago, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		user.MaybePunch(time.Now().In(chicago))
	}
}
