package models

import (
	"backend/ent"
	"backend/ent/user"
	"backend/src/lib"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

/*

	UserModel

*/

type UserModel struct {
	User           *ent.User
	CalendarClient *lib.GoogleCalendarClient
}

func (u UserModel) String() string {
	return fmt.Sprintf("User(name='%s', email='%s', uuid='%s')", u.User.Name, u.User.EmailAddress, u.User.UUID)
}

func (u *UserModel) LoadCalendarClient() bool {
	if u.CalendarClient == nil {
		userCalendarClient, err := lib.NewCalendarClient(u.User.UUID.String())
		if err != nil {
			return false
		}
		u.CalendarClient = userCalendarClient
	}
	return true
}

func UserCreate(name, password, email string) (*UserModel, error) {
	dbClient := lib.DbCtx
	u, err := dbClient.Client.User.
		Create().
		SetName(name).
		SetPassword(hashPassword(password)).
		SetEmailAddress(email).
		Save(dbClient.Context)

	if err != nil {
		return nil, err
	}
	user := UserModel{User: u}
	return &user, nil
}

func UserFind(key string) (*UserModel, error) {
	dbClient := lib.DbCtx
	user_id, err := uuid.Parse(key)
	var users []*ent.User
	if err != nil {
		users, err = dbClient.Client.User.
			Query().
			Where(user.EmailAddress(key)).
			Limit(1).
			All(dbClient.Context)
	} else {
		users, err = dbClient.Client.User.
			Query().
			Where(user.UUID(user_id)).
			Limit(1).
			All(dbClient.Context)
	}

	if err != nil {
		return nil, err
	}

	// Return 404
	if len(users) == 0 {
		return nil, errors.New("User Not Found")
	}

	// Finding first
	var user UserModel
	for _, u := range users {
		user = UserModel{User: u}
		break
	}

	return &user, nil
}

func (u *UserModel) CalendarEventsWithFilters(startDate, endDate string) (lib.DayTimeBlock, error) {
	var calendarEvents lib.DayTimeBlock

	if !u.LoadCalendarClient() {
		return calendarEvents, errors.New("Cannot load Calendar Client")
	}

	// calendarEvents, err := u.CalendarClient.FetchEventsWithFilters(helpers.TimeFormatRFC3339(helpers.TimeToUTC(startDate)), helpers.TimeFormatRFC3339(helpers.TimeToUTC(endDate)))
	calendarEvents, err := u.CalendarClient.FetchEventsWithFilters(startDate, endDate)
	return calendarEvents, err
}

// Returns user events that are already in the calendar
func (u *UserModel) CalendarEvents() (lib.DayTimeBlock, error) {
	var calendarEvents lib.DayTimeBlock

	if !u.LoadCalendarClient() {
		return calendarEvents, errors.New("Cannot load Calendar Client")
	}

	calendarEvents, err := u.CalendarClient.FetchEvents()
	return calendarEvents, err
}

func (u *UserModel) Calendars() ([]string, error) {
	if !u.LoadCalendarClient() {
		return make([]string, 0), errors.New("Cannot load Calendar Client")
	}

	return u.CalendarClient.ListCalendars()
}

func (u *UserModel) AddEvents(events []*Event) bool {
	if !u.LoadCalendarClient() {
		return false
	}

	for _, event := range events {
		for _, calEvent := range event.Confirm() {
			err := u.CalendarClient.AddEvent(calEvent) // TODO: Make sense of the error
			if err != nil {
				panic(err)
			}
		}
	}
	return true
}

func (u UserModel) Events() ([]Event, error) {
	var myEvents []Event
	/*
		Going through all skills, tasks, and other user defined events for the week
	*/
	return myEvents, nil
}

func (u UserModel) Login(password string, bypassCheck bool) (string, error) {
	if !bypassCheck && !u.checkPasswordHash(password) {
		return "", errors.New("Username and Password did not match")
	}
	return lib.GenerateJWT(u.User.UUID.String())
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (u UserModel) checkPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.User.Password), []byte(password))
	return err == nil
}

func (u UserModel) Marshal() UserJSON { // []byte
	user := UserJSON{u.User.Name, u.User.EmailAddress, u.User.CreatedAt}
	// userJSON, _ := json.Marshal(user)
	return user
}

/*

	JwtUser

*/

type JwtUser struct {
	UserId uuid.UUID
	Model  *UserModel
}

func NewJwtUser(user_id string) *JwtUser {
	uuid, _ := uuid.Parse(user_id)
	return &JwtUser{UserId: uuid}
}

// Loads only when neccessary!
func (j *JwtUser) Load() {
	u, _ := UserFind(j.UserId.String())
	j.Model = u
}

/*

	UserJSON

*/

type UserJSON struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
