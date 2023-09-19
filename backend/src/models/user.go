
package models

import (
	"time"
	"backend/ent"
	// "encoding/json"
	"backend/src/lib"
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

func UserCreate(name, password, email string) UserModel {
	dbClient := lib.DbCtx
	u, err := dbClient.Client.User.
				Create().
				SetName(name).
				SetPassword(hashPassword(password)).
				SetEmailAddress(email).
				Save(dbClient.Context)

	user := UserModel{u}
	if err != nil {
		panic(err)
	}
	return user
}

func hashPassword(password string) string {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes)
}

func (u UserModel) checkPasswordHash(hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.user.Password))
    return err == nil
}

func (u UserModel) Marshal() UserJSON { // []byte
	user := UserJSON{u.user.Name, u.user.EmailAddress, u.user.CreatedAt}
	// userJSON, _ := json.Marshal(user)
	return user
}