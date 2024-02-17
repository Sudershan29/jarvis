package models

import (
	"backend/ent"
	"backend/ent/task"
	"backend/ent/user"
	"backend/src/lib"
	"time"
)

type TaskModel struct {
	Task *ent.Task
}

type TaskJSON struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Duration       int       `json:"duration"`
	Deadline       time.Time `json:"deadline"`
	Categories     []string  `json:"categories"`
	TimePreference []string  `json:"timepreference"`
}

func (s TaskModel) Categories() ([]*ent.Category, error) {
	dbClient := lib.DbCtx
	return s.Task.QueryCategories().All(dbClient.Context)
}

func (s TaskModel) TimePreferences() ([]*ent.TimePreference, error) {
	dbClient := lib.DbCtx
	return s.Task.QueryTimePreferences().All(dbClient.Context)
}

func (s TaskModel) Marshal() TaskJSON {
	catArr := make([]string, 0)
	categories, _ := s.Categories()
	for _, cat := range categories {
		catArr = append(catArr, cat.Name)
	}
	timePrefArr := make([]string, 0)
	timePreferences, _ := s.TimePreferences()
	for _, tp := range timePreferences {
		timePrefArr = append(timePrefArr, tp.Day)
	}

	return TaskJSON{s.Task.Name, s.Task.Description, s.Task.Duration, s.Task.Deadline, catArr, timePrefArr}
}

/* * * * * * * * * * * *

		APIs

* * * * * * * * * * * * */

func TaskCreate(name, description string, duration int, deadline time.Time, categories []string, timePreferences []string, currUser *JwtUser) (*TaskModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	sOrm := dbClient.Client.Task.
		Create().
		SetName(name).
		SetDuration(duration).
		SetDeadline(deadline).
		SetUser(currUser.Model.User).
		SetDescription(description)

	for _, cat := range categories {
		catModel, err := CategoryFindOrCreate(cat, currUser)
		if err != nil {
			continue
		}
		sOrm.AddCategories(catModel.Category)
	}

	for _, timePref := range timePreferences {
		timePrefModel, err := TimePreferenceFindOrCreate(timePref)
		if err != nil {
			continue
		}
		sOrm.AddTimePreferences(timePrefModel.TimePreference)
	}

	s, err := sOrm.Save(dbClient.Context)

	if err != nil {
		return nil, err
	}

	task := TaskModel{s}
	return &task, nil
}

func TaskShowAll(currUser *JwtUser) ([]*TaskModel, error) {
	dbClient := lib.DbCtx
	tasks, err := dbClient.Client.Task.
		Query().
		Where(task.HasUserWith(user.UUID(currUser.UserId))).
		All(dbClient.Context)

	if err != nil {
		return make([]*TaskModel, 0), err
	}

	result := make([]*TaskModel, 0)
	for _, task := range tasks {
		result = append(result, &TaskModel{task})
	}
	return result, nil
}

func TaskDelete(taskID int, currUser *JwtUser) error {
	dbClient := lib.DbCtx
	// Find the task by ID and ensure it belongs to the current user
	task, err := dbClient.Client.Task.
		Query().
		Where(task.ID(taskID), task.HasUserWith(user.UUID(currUser.UserId))).
		Only(dbClient.Context)
	if err != nil {
		return err
	}

	// Delete the task
	return dbClient.Client.Task.DeleteOne(task).Exec(dbClient.Context)
}
