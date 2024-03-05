package lib

import (
	"fmt"
	"sort"
	"time"
)

type TimeBlock struct {
	name               string    `json:"name"`
	startTime, endTime time.Time `json:"start_time,omitempty","end_time,omitempty"`
	internal           bool      `json:"internal, omit"` // Sleep, and breakfast timing do not have to be scheduled
	id                 int64     `json:"id"`
}

type DayTimeBlock []*TimeBlock

func (by DayTimeBlock) Less(i, j int) bool { return by[i].Less(by[j]) }

func (day DayTimeBlock) Len() int { return len(day) }

func (day DayTimeBlock) Swap(i, j int) { day[i], day[j] = day[j], day[i] }

func (day DayTimeBlock) MergeOverlaps() DayTimeBlock {
	var mergedEvents DayTimeBlock
	sort.Sort(day)

	if day.Len() > 0 {
		mergedEvents = append(mergedEvents, day[0])
		previous := mergedEvents[0]

		for i := 1; i < day.Len(); i++ {
			if !previous.Merge(*day[i]) { // If you can merge events on overlap, then merge, else make it current
				mergedEvents = append(mergedEvents, day[i])
				previous = day[i]
			}
		}
	}

	return mergedEvents
}

func NewTimeBlock(name string, startTime, endTime time.Time, internal bool, id int64) *TimeBlock {
	return &TimeBlock{name, startTime, endTime, internal, id}
}

func (t TimeBlock) String() string {
	startTimeStr := t.startTime.Format("15:04")
	endTimeStr := t.endTime.Format("15:04")

	return fmt.Sprintf("%s (%s - %s)", t.name, startTimeStr, endTimeStr)
}

func (t1 *TimeBlock) Less(t2 *TimeBlock) bool {
	if t1.startTime.Before(t2.startTime) {
		return true
	} else if t1.startTime.Equal(t2.startTime) {
		return t1.endTime.Before(t2.endTime)
	}
	return false
}

func (t1 *TimeBlock) Merge(t2 TimeBlock) bool {
	// Check for overlap before merging
	if !t1.Overlaps(t2) {
		return false // Don't merge if there's no overlap
	}

	t1.name += ", " + t2.name // Combine names with a comma separator
	t1.startTime = minTime(t1.startTime, t2.startTime)
	t1.endTime = maxTime(t1.endTime, t2.endTime)
	return true
}

func (t1 *TimeBlock) Overlaps(t2 TimeBlock) bool {
	return t1.OverlapInterval(t2.startTime, t2.endTime)
}

func (t1 *TimeBlock) OverlapInterval(start, end time.Time) bool {
	if t1.startTime.After(start) {
		return t1.startTime.Before(end)
	} else {
		return t1.endTime.After(start)
	}
}

func (t1 TimeBlock) StartTime() time.Time {
	return t1.startTime
}

func (t1 TimeBlock) EndTime() time.Time {
	return t1.endTime
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
