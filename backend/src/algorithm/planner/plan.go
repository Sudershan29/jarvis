package planner

import (
	// "fmt"
	"backend/src/helpers"
	"backend/src/lib"
	"backend/src/models"
	"fmt"
	"sort"
	"time"
)

/*

This block will contain the logic behind the scheduling algorithm

1. Creates an imaginary ranges of sleep, lunch, dinner and other essential


*/

func PrepareTimeTable(userUUID string) {
	timezone := "America/Chicago"
	user, _ := models.UserFind(userUUID)
	noOfDays := 3
	userTimezone, _ := time.LoadLocation(timezone)
	tomorrow := time.Now().In(userTimezone).Add(24 * time.Hour)
	startDate := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, userTimezone) // NOTE: Fix this later
	// taskQueues := make([]DayTimeBlock, noOfDays)

	allEvents := Accumulate(userUUID, startDate)

	for day := 0; day < noOfDays; day++ {

		currentDate := startDate.AddDate(0, 0, day)

		calendarEvents, _ := user.CalendarEventsWithFilters(helpers.TimeFormatRFC3339(helpers.TimeToUTC(currentDate)), helpers.TimeFormatRFC3339(helpers.TimeToUTC(currentDate.Add(24*time.Hour))))
		calendarEvents = append(calendarEvents, AddRoutine(timezone, currentDate)...)
		calendarEvents = calendarEvents.MergeOverlaps()

		// You get the tasks as per priority, frequency and schedule them now ( depending on day )
		var events models.CalendarEvent
		// time.ParseDuration("2h")

		sort.Slice(allEvents, func(i, j int) bool {
			return CompareEvents(allEvents[i], allEvents[j], currentDate)
		})

		// events = append(events, models.NewEvent(userUUID, "Leetcode", 90*time.Minute))
		// events = append(events, models.NewEvent(userUUID, "Jobs", 60*time.Minute))
		// events = append(events, models.NewEvent(userUUID, "Grading", 180*time.Minute))
		// events = append(events, models.NewEvent(userUUID, "Energy Disclosure", 60*time.Minute))
		// events = append(events, models.NewEvent(userUUID, "Project", 120*time.Minute))

		breakDuration := 5 * time.Minute
		breathingRoom := 15 * time.Minute
		minWorkTime := 30 * time.Minute
		// allowSplit    := true			// allows one event to be split between events

		// Naive implementation with good breathing room
		nextStart := calendarEvents[0].EndTime()
		idx := 1
		for _, eventWrapper := range allEvents {
			event, err := models.NewEvent(eventWrapper, currentDate)

			if err != nil {
				continue
			}

			durationLeft := event.Length()

			for durationLeft != 0 && idx < len(calendarEvents) {
				if !calendarEvents[idx].OverlapInterval(nextStart, nextStart.Add(event.Length())) {
					event.Schedule(nextStart, nextStart.Add(durationLeft)) // TODO: Add break every 30minutes/1hr
					durationLeft = 0
					nextStart = nextStart.Add(event.Length() + breathingRoom)
				} else if calendarEvents[idx].StartTime().Sub(nextStart) >= minWorkTime { // TODO: Ensure user is fine splitting task
					event.Schedule(nextStart, calendarEvents[idx].StartTime().Truncate(breakDuration))
					durationLeft -= calendarEvents[idx].StartTime().Sub(nextStart)
					nextStart = calendarEvents[idx].EndTime()
					idx += 1

					if durationLeft < 10*time.Minute {
						durationLeft = 0
					} else if durationLeft < 30*time.Minute {
						durationLeft = 30 * time.Minute
					}
				} else {
					nextStart = calendarEvents[idx].EndTime()
					idx += 1
				}
			}
			events = append(events, event)
		}

		user.AddEvents(events)

		fmt.Println(events)
	}
}

func AddRoutine(timezone string, currentDate time.Time) lib.DayTimeBlock {
	var routine lib.DayTimeBlock
	// These are routines
	routine = append(routine, lib.NewTimeBlock("Sleep", currentDate, currentDate.Add(8*time.Hour), true))
	routine = append(routine, lib.NewTimeBlock("Breakfast", currentDate.Add(9*time.Hour), currentDate.Add(10*time.Hour), true))
	routine = append(routine, lib.NewTimeBlock("Lunch", currentDate.Add(13*time.Hour+30*time.Minute), currentDate.Add(14*time.Hour+30*time.Minute), true))
	routine = append(routine, lib.NewTimeBlock("Dinner", currentDate.Add(19*time.Hour+30*time.Minute), currentDate.Add(20*time.Hour+30*time.Minute), true))
	routine = append(routine, lib.NewTimeBlock("Sleep", currentDate.Add(23*time.Hour+30*time.Minute), currentDate.Add(23*time.Hour+59*time.Minute), true))

	return routine
}

func GenerateFakeEvents() lib.DayTimeBlock {
	var fakeEvents lib.DayTimeBlock

	// Get events from Calendar
	fakeEvents = append(fakeEvents, lib.NewTimeBlock("Parallel Office", helpers.ParseTimeWithZone("2024-02-17 13:00:00", "America/Chicago"), helpers.ParseTimeWithZone("2024-02-17 13:45:00", "America/Chicago"), true))
	fakeEvents = append(fakeEvents, lib.NewTimeBlock("Badminton", helpers.ParseTimeWithZone("2024-02-17 18:00:00", "America/Chicago"), helpers.ParseTimeWithZone("2024-02-17 20:45:00", "America/Chicago"), true))

	return fakeEvents
}
