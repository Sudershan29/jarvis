
package jwt_test

import (
	"testing"
	"backend/src/lib"
)


func TestJWTGenerate(t *testing.T) {
	_, err := lib.GenerateJWT("test")
    if err != nil {
        t.Errorf("Cannot generate JWT %s", err.Error())
    }
}

func TestJWTAuthorize(t *testing.T){
	token, _ := lib.GenerateJWT("test")
	user := lib.GetUser(token)
	if user != "test" {
		t.Errorf("Expected `user` to be test, but got %s", user)
	}
}