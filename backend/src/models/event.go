
package models

import (
	"fmt"
	"time"
	"backend/src/helpers"
	"google.golang.org/api/calendar/v3"
)

/*
	Represents a Calendar Event that is an interface to 

	1. Task
	2. Skill
	3. Meeting
	4. Hobby


	Big Idea: One `Event` interface that can be passed along, and Scheduler can use it
*/

type Event struct {
	name, userUUID string
	length time.Duration
	starts, ends []time.Time // easier to add breaks to this using Pomodoro Technique
}

type CalendarEvent []*Event

func NewEvent(userUUID, name string, length time.Duration) *Event {
	return &Event{userUUID: userUUID, name: name, length: length}
}

func (e Event) String() string {
	return fmt.Sprintf("Event(name: %s, start: %v, end: %v)", e.name, e.starts, e.ends)
}

func (e *Event) Length() time.Duration {
	return e.length
}

func (e *Event) Schedule(start, end time.Time) {
	e.starts = append(e.starts, start)
	e.ends   = append(e.ends, end)
}

func (e Event) AsCalendarEvents() []*calendar.Event {
	var result []*calendar.Event
	for idx, _ := range e.starts {
		result = append(result, &calendar.Event{
						Summary: e.name + " [Jarvis]",
						Start: &calendar.EventDateTime{
							DateTime: helpers.TimeFormatRFC3339(helpers.TimeToUTC(e.starts[idx])),
						},
						End: &calendar.EventDateTime{
							DateTime: helpers.TimeFormatRFC3339(helpers.TimeToUTC(e.ends[idx])),
						},
					})
		}
	return result
}