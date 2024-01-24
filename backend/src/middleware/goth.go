package middleware

import (
	"os"
	"encoding/json"
	"path/filepath"
	"github.com/markbates/goth/providers/google"
)

type googleOAuthWeb struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  []string `json:"redirect_uris"`
}

type googleOAuthFile struct {
	Web googleOAuthWeb `json:"web"`
}

func GoogleProvider() *google.Provider {
	absPath, _ := filepath.Abs("config/google_oauth.json")
	byteValue, _ := os.ReadFile(absPath)
	var file googleOAuthFile
	json.Unmarshal(byteValue, &file)
	return google.New(file.Web.ClientId, file.Web.ClientSecret, file.Web.RedirectURI[0])
}