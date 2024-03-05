package planner

import (
	"backend/src/models"
	"time"
)

func Accumulate(userUUID string, date time.Time) []models.CalendarEventAssigner {
	user := models.NewJwtUser(userUUID)
	user.Load()

	var allEvents []models.CalendarEventAssigner
	// Find all skills that are pending, and that can be potentially ordered today

	skills, _ := models.SkillShowAll(user)
	for _, skill := range skills {
		allEvents = append(allEvents, skill)
	}

	// Find all tasks that are pending, and that can be potentially ordered today
	// tasks, _ := models.TaskShowAll(user)
	// for _, task := range tasks {
	// 	allEvents = append(allEvents, task)
	// }

	return allEvents
}
