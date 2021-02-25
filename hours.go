package main

import (
	"time"
)

type Hours struct {
	Name  string `db:"name" json:"name"`
	Hours string `db:"hours" json:"hours"`
}

func GetHours(begin time.Time, end time.Time) ([]Hours, error) {
	var hours []Hours

	err := db.Select(&hours, `
		select
			u.name
			, coalesce(
			sum(julianday("out") -
			julianday("in")) * 24
			, 0
		) as hours
		from punches p
		join users u on p.user_id = u.id
		where "in" between $1 and $2
		group by u.name, date("in")
		order by hours desc
`, begin.Format(time.RFC3339), end.Format(time.RFC3339))

	return hours, err
}
