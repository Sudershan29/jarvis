package models

import (
	"backend/ent"
	"backend/src/lib"
	"backend/ent/skill"
	"backend/ent/user"
)

type SkillModel struct {
	Skill *ent.Skill
}

type SkillJSON struct {
	Name 	   string `json:"name"`
	Level	   string `json:"level"`  
	Duration   int 	 `json:"duration"`
	Categories []string `json:categories`
}

func SkillCreate(name, level string, duration int, categories []string, currUser *JwtUser) (*SkillModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	sOrm := dbClient.Client.Skill.
				Create().
				SetName(name).
				SetUser(currUser.Model.User).
				SetDuration(duration).
				SetLevel(level)

	for _, cat := range categories {
		catModel, err := CategoryFindOrCreate(cat, currUser)
		if err != nil { continue }
		sOrm.AddCategories(catModel.Category)
	}

	s, err := sOrm.Save(dbClient.Context)

	if err != nil { return nil, err }

	skill := SkillModel{s}
	return &skill, nil
}

func SkillShowAll(currUser *JwtUser) ([]*SkillModel, error) {
	dbClient := lib.DbCtx
	skills, err := dbClient.Client.Skill.
					Query().
					Where(skill.HasUserWith(user.UUID(currUser.UserId))).
					All(dbClient.Context)
	
	if err != nil {
		return make([]*SkillModel, 0), err
	}

	result := make([]*SkillModel, 0)
	for _, skill := range skills {
		result = append(result, &SkillModel{skill})
	}
	return result, nil
}

func (s SkillModel) Categories() ([]*ent.Category, error){
	dbClient := lib.DbCtx
	return s.Skill.QueryCategories().All(dbClient.Context)
}

func (s SkillModel) Marshal() SkillJSON {
	catArr := make([]string, 0)
	categories, _ := s.Categories()
	for _, cat := range categories { catArr = append(catArr, cat.Name) }
	return SkillJSON{s.Skill.Name, s.Skill.Level, s.Skill.Duration, catArr}
}