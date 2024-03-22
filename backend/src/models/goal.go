package models

import (
	"backend/ent"
	"backend/ent/goal"
	"backend/ent/user"
	"backend/src/lib"
)

type GoalModel struct {
	Goal *ent.Goal
}

type GoalJSON struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
}

func (s GoalModel) Categories() ([]*ent.Category, error) {
	dbClient := lib.DbCtx
	return s.Goal.QueryCategories().All(dbClient.Context)
}

func (s GoalModel) Marshal() GoalJSON {
	catArr := make([]string, 0)
	categories, _ := s.Categories()
	for _, cat := range categories {
		catArr = append(catArr, cat.Name)
	}
	return GoalJSON{s.Goal.Name, s.Goal.Description, catArr}
}

/* * * * * * * * * * * *

		APIs

* * * * * * * * * * * * */

func GoalCreate(name, description string, categories []string, currUser *JwtUser) (*GoalModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	sOrm := dbClient.Client.Goal.
		Create().
		SetName(name).
		SetUser(currUser.Model.User).
		SetDescription(description)

	for _, cat := range categories {
		catModel, err := CategoryFindOrCreate(cat, currUser)
		if err != nil {
			continue
		}
		sOrm.AddCategories(catModel.Category)
	}

	s, err := sOrm.Save(dbClient.Context)

	if err != nil {
		return nil, err
	}

	goal := GoalModel{s}
	return &goal, nil
}

func GoalShowAll(currUser *JwtUser) ([]*GoalModel, error) {
	dbClient := lib.DbCtx
	goals, err := dbClient.Client.Goal.
		Query().
		Where(goal.HasUserWith(user.UUID(currUser.UserId))).
		All(dbClient.Context)

	if err != nil {
		return make([]*GoalModel, 0), err
	}

	result := make([]*GoalModel, 0)
	for _, goal := range goals {
		result = append(result, &GoalModel{goal})
	}
	return result, nil
}
