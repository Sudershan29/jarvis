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
	event1Preferred := event1.IsPreferredDay(weekday)
	event2Preferred := event2.IsPreferredDay(weekday)
	if event1Preferred && event2Preferred {
		event1HoursLeft := event1.HoursLeft(currDate)
		event2HoursLeft := event2.HoursLeft(currDate)
		if event1HoursLeft == event2HoursLeft {
			return event1.PreferredDaysLeft(currDate) < event2.PreferredDaysLeft(currDate)
		} else {
			return event1HoursLeft > event2HoursLeft
		}
	} else {
		return event1Preferred
	}
}
