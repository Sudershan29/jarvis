package category_test

import (
	"testing"
	// "backend/src/lib"
	"backend/src/models"
	"backend/tests/mocks"
)

// TODO: Fix this!
func init() {
	mocks.SetDatabaseCredentials()
}

func TestCategoryCreate(t *testing.T) {
	_, err := models.CategoryCreate("testing")

	if err != nil {
		t.Errorf("Cannot create category %s", err.Error())
	}
}
