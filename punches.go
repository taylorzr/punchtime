package main

import "time"

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
