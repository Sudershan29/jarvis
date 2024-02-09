
package models

import (
	"fmt"
	"time"
	"errors"
	"backend/ent"
	"backend/src/lib"
	"backend/ent/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

/*

	UserModel

*/

type UserModel struct {
	User *ent.User
}

func (u UserModel) String() string {
	return fmt.Sprintf("User(name='%s', email='%s', uuid='%s')", u.User.Name, u.User.EmailAddress, u.User.UUID)
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
	user := UserModel{u}
	return &user, nil
}

func UserFind(key string) (*UserModel, error) {
	dbClient := lib.DbCtx
	user_id, err := uuid.Parse(key)
	var users []*ent.User
	if err != nil{
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

	if err != nil { return nil, err }

	// Return 404
	if len(users) == 0 { return nil, errors.New("User Not Found") }

	// Finding first
	var user UserModel
	for _, u := range users { user = UserModel{u}; break; }

	return &user, nil
}

// Returns user events that are already in the calendar
func (u UserModel) Calendar() ([]string, error) {
	var calendarEvents []string
	// fmt.Println(u.User.UUID.String())
	// code, err := lib.GetSavedCalendar(u.User.UUID.String())
	userCalendarClient, err := lib.NewCalendarClient(u.User.UUID.String())
	if err != nil {
		return calendarEvents, err;
	}

	calendarEvents, err = userCalendarClient.FetchEvents()
	if err != nil {
		return calendarEvents, err;
	}

	return calendarEvents, nil
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
	Model *UserModel
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
	Name 	  string `json:"name"`
	Email 	  string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}