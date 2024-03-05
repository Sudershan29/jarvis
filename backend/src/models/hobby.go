package models

import (
	"backend/ent"
	"backend/ent/hobby"
	"backend/ent/user"
	"backend/src/lib"
)

type HobbyModel struct {
	Hobby *ent.Hobby
}

type HobbyJSON struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
}

func (s HobbyModel) Categories() ([]*ent.Category, error) {
	dbClient := lib.DbCtx
	return s.Hobby.QueryCategories().All(dbClient.Context)
}

func (s HobbyModel) Marshal() HobbyJSON {
	catArr := make([]string, 0)
	categories, _ := s.Categories()
	for _, cat := range categories {
		catArr = append(catArr, cat.Name)
	}
	return HobbyJSON{s.Hobby.Name, s.Hobby.Description, catArr}
}

/* * * * * * * * * * * *

		APIs

* * * * * * * * * * * * */

func HobbyCreate(name, description string, categories []string, currUser *JwtUser) (*HobbyModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	sOrm := dbClient.Client.Hobby.
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

	hobby := HobbyModel{s}
	return &hobby, nil
}

func HobbyShowAll(currUser *JwtUser) ([]*HobbyModel, error) {
	dbClient := lib.DbCtx
	hobbies, err := dbClient.Client.Hobby.
		Query().
		Where(hobby.HasUserWith(user.UUID(currUser.UserId))).
		All(dbClient.Context)

	if err != nil {
		return make([]*HobbyModel, 0), err
	}

	result := make([]*HobbyModel, 0)
	for _, hobby := range hobbies {
		result = append(result, &HobbyModel{hobby})
	}
	return result, nil
}
