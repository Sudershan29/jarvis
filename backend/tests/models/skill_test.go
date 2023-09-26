
package skill_test

import (
	"testing"
	"backend/src/models"
)

func TestSkillCreate(t *testing.T) {
	skillORM, err := models.SkillCreate("testing", "beginner", 10, []string{"test1", "test2"})

	if err != nil {
        t.Errorf("Cannot create category %s", err.Error())
    }

	categories, err := skillORM.Categories()

	if err != nil || len(categories) != 2 {
		t.Errorf("Expected 2 categories but got %d", len(categories))
	}
}

