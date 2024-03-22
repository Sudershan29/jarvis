package lib

import (
	"sort"
	"time"
)

type GCalendarEvent struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Internal  bool      `json:"internal"` // Sleep, and breakfast timing do not have to be scheduled
	Id        int64     `json:"id"`
}

type GCalendarEventsGroup []*GCalendarEvent

func (by GCalendarEventsGroup) Less(i, j int) bool { return by[i].Less(by[j]) }

func (day GCalendarEventsGroup) Len() int { return len(day) }

func (day GCalendarEventsGroup) Swap(i, j int) { day[i], day[j] = day[j], day[i] }

func (day GCalendarEventsGroup) MergeOverlaps() GCalendarEventsGroup {
	var mergedEvents GCalendarEventsGroup
	sort.Sort(day)

	if day.Len() > 0 {
		mergedEvents = append(mergedEvents, day[0])
		previous := mergedEvents[0]

		for i := 1; i < day.Len(); i++ {
			if !previous.Merge(day[i]) { // If you can merge events on overlap, then merge, else make it current
				mergedEvents = append(mergedEvents, day[i])
				previous = day[i]
			}
		}
	}

	return mergedEvents
}

func NewGCalendarEvent(name string, startTime, endTime time.Time, internal bool, id int64) *GCalendarEvent {
	return &GCalendarEvent{name, startTime, endTime, internal, id}
}

// func (t GCalendarEvent) String() string {
// 	startTimeStr := t.startTime.Format("15:04")
// 	endTimeStr := t.endTime.Format("15:04")

// 	return fmt.Sprintf("%s (%s - %s)", t.name, startTimeStr, endTimeStr)
// }

func (t1 *GCalendarEvent) Less(t2 *GCalendarEvent) bool {
	if t1.StartTime.Before(t2.StartTime) {
		return true
	} else if t1.StartTime.Equal(t2.StartTime) {
		return t1.EndTime.Before(t2.EndTime)
	}
	return false
}

func (t1 *GCalendarEvent) Merge(t2 *GCalendarEvent) bool {
	// Check for overlap before merging
	if !t1.Overlaps(t2) {
		return false // Don't merge if there's no overlap
	}

	t1.Name += ", " + t2.Name // Combine names with a comma separator
	t1.StartTime = minTime(t1.StartTime, t2.StartTime)
	t1.EndTime = maxTime(t1.EndTime, t2.EndTime)
	return true
}

func (t1 *GCalendarEvent) Overlaps(t2 *GCalendarEvent) bool {
	return t1.OverlapInterval(t2.StartTime, t2.EndTime)
}

func (t1 *GCalendarEvent) OverlapInterval(start, end time.Time) bool {
	if t1.StartTime.After(start) {
		return t1.StartTime.Before(end)
	} else {
		return t1.EndTime.After(start)
	}
}

// Helper functions for clarity and potential reusability
func minTime(t1, t2 time.Time) time.Time {
	if t1.Before(t2) {
		return t1
	}
	return t2
}

func maxTime(t1, t2 time.Time) time.Time {
	if t1.After(t2) {
		return t1
	}
	return t2
}
