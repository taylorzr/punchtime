package main

import (
	"time"
)

type Hours struct {
	Name       string  `db:"name" json:"name"`
	Hours      float32 `db:"hours" json:"hours"`
	PunchCount int     `db:"punch_count" json:"punch_count"`
}

func GetHours(begin time.Time, end time.Time, windowBegin time.Time, windowEnd time.Time) ([]Hours, error) {
	var hours []Hours

	// FIXME: So this now handles punches where someone is punched in before 9 and after 5,
	// but it doesn't handle punched in for days at time
	// Don't really care, but to fix we should force punch out / in at midnight / whatever we consider
	// beginning of day (5am?)
	// E.g. See March 16th
	err := config.DB.Select(&hours, `
		select
			u.name
			, round(
					coalesce(
						sum(
							julianday(min(datetime(coalesce("out", 'now')), datetime($2)))
							- julianday(max(datetime("in"), datetime($1)))
						) * 24
					, 0)
				, 2) as hours
			, count(p.id) as punch_count
		from punches p
		join users u on p.user_id = u.id
		where datetime("in") between datetime($1) and datetime($2)
		or datetime(coalesce("out", 'now')) between datetime($1) and datetime($2)
		group by u.name

		union

		select
			u.name
			, round(
					(
						julianday(min(coalesce("out", datetime('now')), datetime($2))) -
						julianday(max("in", $1))
					) * 24
				, 2) as hours
			, count(p.id) as punch_count
		from punches p
		join users u on p.user_id = u.id
		where u.name = 'Zach Taylor'
		and datetime("in") between datetime($3) and datetime($4)
		and datetime("in") < datetime($1)
		and datetime("out") > datetime($2)
		group by u.name

	`, begin.Format(time.RFC3339), end.Format(time.RFC3339), windowBegin.Format(time.RFC3339), windowEnd.Format(time.RFC3339))

	return hours, err
}

type FirstLast struct {
	Name  string `db:"name" json:"name"`
	First string `db:"first" json:"first"`
	Last  string `db:"last" json:"last"`
}

func GetFirstLasts(begin time.Time, end time.Time) ([]FirstLast, error) {
	var firstLasts []FirstLast

	err := config.DB.Select(&firstLasts, `
		select u.name, min("in") as first, coalesce(max("out"), '-')  as last from punches p
		join users u on u.id = p.user_id
		where "in" between $1 and $2
		group by u.name
		order by min("in") asc
		;
	`, begin.Format(time.RFC3339), end.Format(time.RFC3339))

	// FIXME: Scanner for queries that returns times
	// for _, fl := range firstLasts {
	// 	fl.First = fl.First.In(config.Timezone)
	// 	fl.Last = fl.Last.In(config.Timezone)
	// }

	return firstLasts, err
}
