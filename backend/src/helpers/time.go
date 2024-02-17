package helpers

import (
	"time"
)

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
