package models

import (
	"backend/ent"
	"backend/ent/calendar"
	"backend/ent/user"
	"backend/src/lib"
)

type CalendarModel struct {
	Calendar *ent.Calendar
}

type CalendarJSON struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	token string `json:"-"`
}

func (s CalendarModel) Marshal() CalendarJSON {
	return CalendarJSON{s.Calendar.Name, s.Calendar.Type, s.Calendar.Token}
}

/* * * * * * * * * * * *

		APIs

* * * * * * * * * * * * */

func CalendarCreate(name, calendarType, token string, currUser *JwtUser) (*CalendarModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	sOrm := dbClient.Client.Calendar.
		Create().
		SetName(name).
		SetType(calendarType).
		SetToken(token).
		SetUser(currUser.Model.User)

	s, err := sOrm.Save(dbClient.Context)

	if err != nil {
		return nil, err
	}

	calendar := CalendarModel{s}
	return &calendar, nil
}

func CalendarShowAll(currUser *JwtUser) ([]*CalendarModel, error) {
	dbClient := lib.DbCtx
	calendars, err := dbClient.Client.Calendar.
		Query().
		Where(calendar.HasUserWith(user.UUID(currUser.UserId))).
		All(dbClient.Context)

	if err != nil {
		return make([]*CalendarModel, 0), err
	}

	result := make([]*CalendarModel, 0)
	for _, calendar := range calendars {
		result = append(result, &CalendarModel{calendar})
	}

	return result, nil
}

func (s *CalendarModel) CalendarUpdateToken(token string) error {
	_, err := s.Calendar.Update().SetToken(token).Save(lib.DbCtx.Context)
	return err
}

func CalendarShow(id int, user *JwtUser) (lib.DayTimeBlock, error) {
	userCalendarClient, err := lib.NewCalendarClient(user.UserId.String())

	if err != nil {
		return lib.DayTimeBlock{}, err
	}

	return userCalendarClient.FetchEvents()
}

func CalendarFindOrCreate(name, calendarType, token string, currUser *JwtUser) (*CalendarModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	sOrm, err := dbClient.Client.Calendar.
		Query().
		Where(calendar.Name(name)).
		Where(calendar.Type(calendarType)).
		Where(calendar.HasUserWith(user.UUID(currUser.UserId))).
		Only(dbClient.Context)

	if err == nil {
		calendar := CalendarModel{sOrm}
		calendar.CalendarUpdateToken(token)
		return &calendar, nil
	}

	return CalendarCreate(name, calendarType, token, currUser)
}
