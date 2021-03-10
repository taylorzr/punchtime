package main

import (
	"fmt"
	"log"
	"time"
)

type DayLog struct {
	Begin time.Time     `json:"begin"`
	End   time.Time     `json:"end"`
	Hours []*UserDayLog `json:"hours"`
}

type UserDayLog struct {
	Name           string  `db:"name" json:"name"`
	Hours          float32 `db:"hours" json:"hours"`
	CoreHours      float32 `db:"core_hours" json:"core_hours"`
	PunchCount     int     `db:"punch_count" json:"punch_count"`
	CorePunchCount int     `db:"core_punch_count" json:"core_punch_count"`
}

func GetDayLog(now time.Time) DayLog {
	begin := beginDay(now)
	end := endDay(now)

	fmt.Printf("Getting hours between %s and %s\n", begin.Format(time.RFC3339), end.Format(time.RFC3339))
	allHours, err := GetHours(begin, end)
	// FIXME: Write http 500?
	if err != nil {
		log.Fatal(err)
	}

	coreBegin := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, now.Location()).UTC()
	coreEnd := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location()).UTC()
	fmt.Printf("Getting hours between %s and %s\n", coreBegin.Format(time.RFC3339), coreEnd.Format(time.RFC3339))
	coreHours, err := GetHours(coreBegin, coreEnd)
	if err != nil {
		log.Fatal(err)
	}

	usersDayLogs := map[string]*UserDayLog{}
	for _, hours := range allHours {
		usersDayLogs[hours.Name] = &UserDayLog{
			Name:       hours.Name,
			Hours:      hours.Hours,
			PunchCount: hours.PunchCount,
		}
	}
	for _, hours := range coreHours {
		dayLog := usersDayLogs[hours.Name]
		dayLog.CoreHours = hours.Hours
		dayLog.CorePunchCount = hours.PunchCount
	}

	index := 0
	var hours = make([]*UserDayLog, len(usersDayLogs))
	for _, usersDayLog := range usersDayLogs {
		hours[index] = usersDayLog
		index++
	}

	return DayLog{
		Begin: begin.In(config.Timezone),
		End:   end.In(config.Timezone),
		Hours: hours,
	}
}

func now() time.Time {
	return time.Now().In(config.Timezone)
}

func beginDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UTC()
}

func endDay(t time.Time) time.Time {
	return beginDay(t).AddDate(0, 0, 1)
}
