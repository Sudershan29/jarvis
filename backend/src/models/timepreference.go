package models

import (
	"backend/ent"
	"backend/ent/timepreference"
	"backend/src/lib"
)

type TimePreferenceModel struct {
	TimePreference *ent.TimePreference
}

type TimePreferenceJSON struct {
	Day string `json:"day"`
}

func TimePreferenceFindOrCreate(day string) (*TimePreferenceModel, error) {
	dbClient := lib.DbCtx
	tp, err := dbClient.Client.TimePreference.
		Query().
		Where(timepreference.Day(day)).
		Only(dbClient.Context)

	if err != nil {
		// If not found, create a new TimePreference
		tp, err = dbClient.Client.TimePreference.
			Create().
			SetDay(day).
			Save(dbClient.Context)
		if err != nil {
			return nil, err
		}
	}

	return &TimePreferenceModel{TimePreference: tp}, nil
}

func TimePreferenceShowAll() ([]*TimePreferenceModel, error) {
	dbClient := lib.DbCtx
	timePreferences, err := dbClient.Client.TimePreference.
		Query().
		All(dbClient.Context)

	if err != nil {
		return make([]*TimePreferenceModel, 0), err
	}

	result := make([]*TimePreferenceModel, 0)
	for _, tp := range timePreferences {
		result = append(result, &TimePreferenceModel{TimePreference: tp})
	}
	return result, nil
}
