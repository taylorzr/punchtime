package main

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
)

func GetUsers() ([]User, error) {
	var users []User

	err := db.Select(&users, "select * from users")

	return users, err
}

type User struct {
	ID      int    `db:"id"`
	SlackID string `db:"slack_id"`
	Name    string `db:"name"`
}

// TODO:
func (user User) Punches() ([]Punch, error) {
	return nil, nil
}

func (user User) LastPunch() (*Punch, error) {
	var punches []Punch

	err := db.Select(&punches, `
		select p.* from punches p
		join users u on u.id = p.user_id
		where u.id = $1
		order by id
		desc limit 1
	`, user.ID)

	if len(punches) == 0 {
		return nil, nil
	}

	return &punches[0], err
}

// curl -sH "Authorization: Bearer $SLACK_TOKEN" 'https://slack.com/api/users.getPresence' | jq .
func (user User) Presence() (*slack.UserPresence, error) {
	// FIXME: If you use a bad id this doesn't return an error just fails silently?!?!
	// E.g. my id is W0162LHTW2C but say you mess up copy/paste and use W0163LHTW2C instead
	presence, err := slackClient.GetUserPresence(user.SlackID)
	return presence, err
}

func (user User) CreatePunch(t time.Time) error {
	_, err := db.Exec(`INSERT INTO punches (user_id, "in") VALUES ($1, $2)`, user.ID, t.Format(time.RFC3339))
	return err
}

func (user User) MaybePunch(t time.Time) error {
	lastPunch, err := user.LastPunch()
	if err != nil {
		return err
	}

	presence, err := user.Presence()
	if err != nil {
		return err
	}

	// Handle last punch null / first punch
	if lastPunch == nil {
		if presence.Presence == "active" {
			fmt.Printf("%s First punch ever\n", user.Name)
			err := user.CreatePunch(t)
			if err != nil {
				return err
			}
		}
		return nil
	}

	// Guaranteed to have a last punch
	if presence.Presence == "active" {
		if lastPunch.Out != nil {
			fmt.Printf("%s punched in\n", user.Name)
			err := user.CreatePunch(t)
			if err != nil {
				return err
			}
		} else {
			// TODO: Put behind verbose flag or something
			fmt.Printf("%s is still active\n", user.Name)
		}
	} else {
		if lastPunch.Out == nil {
			fmt.Printf("%s punched out\n", user.Name)
			err := lastPunch.PunchOut(t)
			if err != nil {
				return err
			}
		} else {
			// TODO: Put behind verbose flag or something
			fmt.Printf("%s is still away\n", user.Name)
		}
	}

	return nil
}
