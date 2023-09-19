
package models

import (
	"time"
	"errors"
	"backend/ent"
	"backend/src/lib"
	"backend/ent/user"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	user *ent.User
}

type UserJSON struct {
	Name 	  string `json:"name"`
	Email 	  string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
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

func UserFind(email string) (*UserModel, error) {
	dbClient := lib.DbCtx
	users, err := dbClient.Client.User.
				Query().
				Where(user.EmailAddress(email)).
				Limit(1).
				All(dbClient.Context)

	if err != nil { return nil, err }

	// Return 404
	if len(users) == 0 { return nil, errors.New("User Not Found") }

	// Finding first
	var user UserModel
	for _, u := range users { user = UserModel{u}; break; }

	return &user, nil
}

func (u UserModel) Login(password string) (string, error) {
	if !u.checkPasswordHash(password) {
		return "", errors.New("Username and Password did not match")
	}
	return lib.GenerateJWT(u.user.Uuid)
}

func hashPassword(password string) string {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes)
}

func (u UserModel) checkPasswordHash(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.user.Password), []byte(password))
    return err == nil
}

func (u UserModel) Marshal() UserJSON { // []byte
	user := UserJSON{u.user.Name, u.user.EmailAddress, u.user.CreatedAt}
	// userJSON, _ := json.Marshal(user)
	return user
}