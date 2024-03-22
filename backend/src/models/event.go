package models

import (
	"backend/ent/proposal"
	"backend/src/helpers"
	"errors"
	"fmt"
	"time"

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

type CalendarEventAssigner interface {
	PreferredDaysLeft(time.Time) int
	HoursLeft(time.Time) int
	PreferredMaxHours() int
	PreferredMinHours() int
	IsPreferredDay(string) bool
	String() string
	AddProposal(proposal *ProposalModel) bool
	Name() string
	HoursNeeded(time.Time) float64
	UserUUID() string
}

type Event struct {
	name, userUUID string
	length         time.Duration
	starts, ends   []time.Time // easier to add breaks to this using Pomodoro Technique
	event          CalendarEventAssigner
}

type CalendarEvent []*Event

func NewEvent(event CalendarEventAssigner, currDay time.Time) (*Event, error) {
	if event.HoursNeeded(currDay) <= 0 {
		return nil, errors.New("does not need hours")
	}

	return &Event{userUUID: event.UserUUID(), name: event.Name(), length: time.Duration(event.HoursNeeded(currDay)) * time.Hour, event: event}, nil
}

func (e Event) String() string {
	return fmt.Sprintf("Event(name: %s, start: %v, end: %v)", e.name, e.starts, e.ends)
}

func (e *Event) Length() time.Duration {
	return e.length
}

func (e *Event) Schedule(start, end time.Time) {
	e.starts = append(e.starts, start)
	e.ends = append(e.ends, end)
}

func (e Event) AsCalendarEvents() []*calendar.Event {
	var result []*calendar.Event
	for idx := range e.starts {
		result = append(result, &calendar.Event{
			Summary: e.ModifiedName(),
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

func (e Event) Confirm() []*calendar.Event {
	var calendarEvents []*calendar.Event

	for idx := range e.starts {
		duration := int(e.ends[idx].Sub(e.starts[idx]).Minutes())
		proposal, err := ProposalCreate(e.ModifiedName(), duration, 0, proposal.StatusPending.String(), e.starts[idx])
		if err != nil {
			panic(err)
		}
		e.event.AddProposal(proposal)
		// TODO: Ensure no failure

		calendarEvents = append(calendarEvents, &calendar.Event{
			Summary: e.ModifiedName(),
			Start: &calendar.EventDateTime{
				DateTime: helpers.TimeFormatRFC3339(helpers.TimeToUTC(e.starts[idx])),
			},
			End: &calendar.EventDateTime{
				DateTime: helpers.TimeFormatRFC3339(helpers.TimeToUTC(e.ends[idx])),
			},
			Id: helpers.IdToCalendarId(int64(proposal.Proposal.ID)),
		})
	}

	return calendarEvents
}

func (e Event) ModifiedName() string {
	return e.name + " [Jarvis]"
}
