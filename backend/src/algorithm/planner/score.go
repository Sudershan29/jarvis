package planner

import (
	"backend/src/models"
	"time"
)

// This module is reposible for scoring the priority of a proposal
// The priority is based on the following factors:
// 1. The number of hours vs number of opportunities to be scheduled
// 2. The number of total hours left for the proposal
// 3. .....

// TODO: Move this to a weighted scoring method
func Score(event models.CalendarEventAssigner, currDate time.Time) int {
	if event.HoursLeft(currDate) == 0 {
		return 0
	}

	return event.HoursLeft(currDate)
}

func CompareEvents(event1, event2 models.CalendarEventAssigner, currDate time.Time) bool {
	weekday := currDate.Weekday().String()
	if event1.IsPreferredDay(weekday) && event2.IsPreferredDay(weekday) {
		if event1.HoursLeft(currDate) == event2.HoursLeft(currDate) {
			return event1.PreferredDaysLeft(currDate) < event2.PreferredDaysLeft(currDate)
		} else {
			return event1.HoursLeft(currDate) > event2.HoursLeft(currDate)
		}
	} else {
		return event1.IsPreferredDay(weekday)
	}
}
