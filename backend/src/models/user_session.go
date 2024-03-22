package models

import (
	"time"
	"github.com/markbates/goth"
)

type UserSessionJSON struct {
	Name 	     string `json:"name"`
	Email 	     string `json:"email"`
	Provider     string `json:"provider"`
	AccessToken  string `json:"-"`
	RefreshToken string `json:"-"`
	ExpiresAt    time.Time `json:"expiry"`
}

func ConvertToJSON(user goth.User) UserSessionJSON {
	return UserSessionJSON{ user.FirstName + " " + user.LastName,
							user.Email, user.Provider, user.AccessToken,
							user.RefreshToken, user.ExpiresAt }
}