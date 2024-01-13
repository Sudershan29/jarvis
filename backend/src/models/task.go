package models

import (
	"time"
	"backend/ent"
	"backend/src/lib"
	"backend/ent/task"
	"backend/ent/user"
)

type TaskModel struct {
	Task *ent.Task
}

type TaskJSON struct {
	Name 	     string    `json:"name"`
	Description	 string    `json:"description"`
	Duration	 int       `json:"duration"`
	Deadline 	 time.Time `json:"deadline"`
	Categories []string    `json:categories`
}

func (s TaskModel) Categories() ([]*ent.Category, error){
	dbClient := lib.DbCtx
	return s.Task.QueryCategories().All(dbClient.Context)
}

func (s TaskModel) Marshal() TaskJSON {
	catArr := make([]string, 0)
	categories, _ := s.Categories()
	for _, cat := range categories { catArr = append(catArr, cat.Name) }
	return TaskJSON{s.Task.Name, s.Task.Description, s.Task.Duration, s.Task.Deadline, catArr}
}

/* * * * * * * * * * * * 

		APIs

* * * * * * * * * * * * */

func TaskCreate(name, description string, duration int, deadline time.Time, categories []string, currUser *JwtUser) (*TaskModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	sOrm := dbClient.Client.Task.
				Create().
				SetName(name).
				SetDuration(duration).
				SetDeadline(deadline).
				SetDescription(description)

	for _, cat := range categories {
		catModel, err := CategoryFindOrCreate(cat, currUser)
		if err != nil { continue }
		sOrm.AddCategories(catModel.Category)
	}

	s, err := sOrm.Save(dbClient.Context)

	if err != nil { return nil, err }

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