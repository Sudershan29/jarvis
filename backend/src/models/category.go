package models

import (
	"strings"
	"backend/ent"
	"backend/src/lib"
	"backend/ent/category"
	"backend/ent/user"
)

type CategoryModel struct {
	Category *ent.Category
}

type CategoryJSON struct {
	Name 	  string `json:"name"`
}

func formatName(name string) string {
	return strings.Title(strings.ToLower(name))
}

func CategoryCreate(name string, currUser *JwtUser) (*CategoryModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	formattedName := formatName(name)
	c, err := dbClient.Client.Category.
				Create().
				SetUser(currUser.Model.User).
				SetName(formattedName).
				Save(dbClient.Context)

	if err != nil {
		return nil, err
	}

	category := CategoryModel{c}
	return &category, nil
}

func CategoryFindOrCreate(name string, currUser *JwtUser) (*CategoryModel, error) {
	dbClient := lib.DbCtx
	formattedName := formatName(name)
	c, err := dbClient.Client.Category.
						Query().
						Where(category.Name(formattedName),
							  category.HasUserWith(user.UUID(currUser.UserId))).
						Limit(1).
						All(dbClient.Context)
	if err != nil || len(c) != 1 {
		return CategoryCreate(formattedName, currUser)
	}

	category := CategoryModel{c[0]}
	return &category, nil
}