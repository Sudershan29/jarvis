package planner

import (
	// "fmt"
	"backend/src/helpers"
	"backend/src/lib"
	"backend/src/models"
	"errors"
	"fmt"
	"sort"
	"time"
)

/*

This block will contain the logic behind the scheduling algorithm

1. Creates an imaginary ranges of sleep, lunch, dinner and other essential


*/

// func PrepareTimeTable(userUUID string) {
// 	user, _ := models.UserFind(userUUID)
// 	timezone := user.TimeZone()
// 	noOfDays := 3
// 	userTimezone, _ := time.LoadLocation(timezone)
// 	tomorrow := time.Now().In(userTimezone).Add(24 * time.Hour)
// 	startDate := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, userTimezone)
// 	allEvents := PrepareAllActivities(userUUID, startDate)

// 	for day := 0; day < noOfDays; day++ {
// 		currentDate := startDate.AddDate(0, 0, day)
// 		PlanADay("primary", currentDate, user, allEvents)
// 	}
// }

func PrepareTimeTable(userUUID string, days int) error {
	user, err := models.UserFind(userUUID)
	if err != nil {
		return err
	}

	if !user.LoadCalendarClient() {
		return errors.New("cannot load Calendar Client")
	}

	// Setting tomorrow and day
	userTimezone, _ := time.LoadLocation(user.TimeZone())
	tomorrow := time.Now().In(userTimezone).Add(24 * time.Hour)
	startDate := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, userTimezone)

	// Clear all Jarvis events from those days
	jarvisCalendarId, err := user.CalendarClient.FindOrCreateJarvisCalendar(userUUID, user.TimeZone())
	if err != nil {
		return err
	}

	// Clearing Jarvis events, and marking them as cancelled in the database
	for day := 0; day < days; day++ {
		currentDate := startDate.AddDate(0, 0, day)
		user.CalendarClient.ClearEventsInCalendarByDate(jarvisCalendarId, currentDate.Format("2017-09-07"))
	}

	// Get all events
	allEvents := PrepareAllActivities(userUUID, startDate)
	for day := 0; day < days; day++ {
		currentDate := startDate.AddDate(0, 0, day)
		PlanADay(jarvisCalendarId, currentDate, user, allEvents)
	}

	return nil
}

func PlanADay(calendarId string, currentDate time.Time, user *models.UserModel, allEvents []models.CalendarEventAssigner) {
	timezone := user.TimeZone()
	calendarEvents, _ := user.CalendarEventsWithFilters(helpers.TimeFormatRFC3339(helpers.TimeToUTC(currentDate)),
		helpers.TimeFormatRFC3339(helpers.TimeToUTC(currentDate.Add(24*time.Hour))))

	calendarEvents = append(calendarEvents, AddRoutine(timezone, currentDate)...)
	calendarEvents = calendarEvents.MergeOverlaps()

	// You get the tasks as per priority, frequency and schedule them now ( depending on day )
	var events models.CalendarEvent

	sort.Slice(allEvents, func(i, j int) bool {
		return CompareEvents(allEvents[i], allEvents[j], currentDate)
	})

	breakDuration := 5 * time.Minute
	breathingRoom := 15 * time.Minute
	minWorkTime := 30 * time.Minute
	// allowSplit    := true			// allows one event to be split between events

	// Naive implementation with good breathing room
	nextStart := calendarEvents[0].EndTime
	idx := 1
	for _, eventWrapper := range allEvents {
		event, err := models.NewEvent(eventWrapper, currentDate)

		if err != nil {
			continue
		}

		durationLeft := event.Length()

		for durationLeft != 0 && idx < len(calendarEvents) {
			// If there is no overlap
			if !calendarEvents[idx].OverlapInterval(nextStart, nextStart.Add(event.Length())) {
				event.Schedule(nextStart, nextStart.Add(durationLeft)) // TODO: Add break every 30minutes/1hr
				durationLeft = 0
				nextStart = nextStart.Add(event.Length() + breathingRoom)
			} else if calendarEvents[idx].StartTime.Sub(nextStart) >= minWorkTime { // TODO: Ensure user is fine splitting task
				event.Schedule(nextStart, calendarEvents[idx].StartTime.Truncate(breakDuration))
				durationLeft -= calendarEvents[idx].StartTime.Sub(nextStart)
				nextStart = calendarEvents[idx].EndTime
				idx += 1

				if durationLeft < 10*time.Minute {
					durationLeft = 0
				} else if durationLeft < 30*time.Minute {
					durationLeft = 30 * time.Minute
				}
			} else {
				nextStart = calendarEvents[idx].EndTime
				idx += 1
			}
		}
		events = append(events, event)
	}

	user.AddEvents(calendarId, events)

	fmt.Println(events)
}

func AddRoutine(timezone string, currentDate time.Time) lib.GCalendarEventsGroup {
	var routine lib.GCalendarEventsGroup
	// These are routines
	routine = append(routine, lib.NewGCalendarEvent("Sleep", currentDate, currentDate.Add(8*time.Hour), true, -1))
	routine = append(routine, lib.NewGCalendarEvent("Breakfast", currentDate.Add(9*time.Hour), currentDate.Add(10*time.Hour), true, -1))
	routine = append(routine, lib.NewGCalendarEvent("Lunch", currentDate.Add(13*time.Hour+30*time.Minute), currentDate.Add(14*time.Hour+30*time.Minute), true, -1))
	routine = append(routine, lib.NewGCalendarEvent("Dinner", currentDate.Add(19*time.Hour+30*time.Minute), currentDate.Add(20*time.Hour+30*time.Minute), true, -1))
	routine = append(routine, lib.NewGCalendarEvent("Sleep", currentDate.Add(23*time.Hour+30*time.Minute), currentDate.Add(23*time.Hour+59*time.Minute), true, -1))

	return routine
}

func GenerateFakeEvents() lib.GCalendarEventsGroup {
	var fakeEvents lib.GCalendarEventsGroup

	// Get events from Calendar
	fakeEvents = append(fakeEvents, lib.NewGCalendarEvent("Parallel Office", helpers.ParseTimeWithZone("2024-02-17 13:00:00", "America/Chicago"), helpers.ParseTimeWithZone("2024-02-17 13:45:00", "America/Chicago"), true, -1))
	fakeEvents = append(fakeEvents, lib.NewGCalendarEvent("Badminton", helpers.ParseTimeWithZone("2024-02-17 18:00:00", "America/Chicago"), helpers.ParseTimeWithZone("2024-02-17 20:45:00", "America/Chicago"), true, -1))

	return fakeEvents
}
