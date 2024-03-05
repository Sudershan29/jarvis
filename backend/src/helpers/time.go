package helpers

import (
	"strings"
	"time"
)

var WEEK = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
var WEEKNUMBER = map[string]int{
	"Sunday":    0,
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
}

func ParseTimeWithZone(t, timezone string) time.Time {
	// Central Time location
	ctLoc, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}

	ctTime, err := time.ParseInLocation("2006-01-02 15:04:05", t, ctLoc)
	if err != nil {
		panic(err)
	}

	return ctTime
}

func TimeToUTC(otherTimezone time.Time) time.Time {
	return otherTimezone.In(time.UTC)
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TimeFormatRFC3339(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z")
}

func ConvertToTimezone(t time.Time, timezone string) time.Time {
	ctLoc, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}

	return t.In(ctLoc)
}

// Assuming week goes from Sunday to Saturday

func StartOfWeek(t time.Time) time.Time {
	daysAgo := int(t.Weekday())
	startOfWeek := t.AddDate(0, 0, -daysAgo)
	return time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, startOfWeek.Location())
}

func EndOfWeek(t time.Time) time.Time {
	daysAhead := 6 - int(t.Weekday())
	endOfWeek := t.AddDate(0, 0, daysAhead)
	return time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 23, 59, 59, 0, endOfWeek.Location())
}

func NextDays(currentDay time.Time) []string {
	currIdx := int(currentDay.Weekday())
	return WEEK[currIdx:]
}

func NumberOfDaysLeftThisWeek(currDay time.Time, days []string) int {
	todayIdx := int(currDay.Weekday())
	count := 0
	for _, day := range days {
		if WEEKNUMBER[strings.Title(strings.ToLower(day))] >= todayIdx {
			count++
		}
	}
	return count
}

func NumberOfPreferredDaysBetween(currDate, nextDate time.Time, preferredDays []string) int {
	weeks := int(nextDate.Sub(currDate).Hours() / 24 / 7)

	return NumberOfDaysLeftThisWeek(currDate, preferredDays) + len(preferredDays)*weeks
}
