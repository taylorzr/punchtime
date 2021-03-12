package main

import (
	"log"
	"time"
)

// FIXME: Rename this and normal Punch
type PPunch struct {
	Name string `db:"name" json:"name"`
	ID   int    `db:"id" json:"id"`
	In   string `db:"in" json:"in"`
	Out  string `db:"out" json:"out"`
}

func PunchesFor(username string) []PPunch {
	var punches []PPunch

	err := config.DB.Select(&punches, `
		select
		  u.name
			, p.id
			, p."in"
			, coalesce(p."out", '-') as out
		from punches p
		join users u on u.id = p.user_id
		where lower(u.name) = lower($1)
		order by p.id desc
	`, username)

	if err != nil {
		log.Fatal(err)
	}

	return punches

}

// TODO: Map time string in sqlite to actual time in golang
// https://stackoverflow.com/questions/41375563/unsupported-scan-storing-driver-value-type-uint8-into-type-string
// https://gist.github.com/jmoiron/6979540
type Punch struct {
	ID       int     `db:"id"`
	UserID   int     `db:"user_id"`
	In       string  `db:"in"`
	Out      *string `db:"out"`
	Presence string  `db:"presence"`
}

func (punch Punch) PunchOut(t time.Time) error {
	_, err := config.DB.Exec(`
		update punches
		set "out" = $1
		where id = $2
	`, t.Format(time.RFC3339), punch.ID)
	// ^ FIXME: Should probably use chicago timezone
	return err
}
