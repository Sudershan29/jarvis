// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/preference"
	"backend/ent/schema"
	"backend/ent/skill"
	"backend/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	preferenceFields := schema.Preference{}.Fields()
	_ = preferenceFields
	// preferenceDescFreeWeekends is the schema descriptor for free_weekends field.
	preferenceDescFreeWeekends := preferenceFields[0].Descriptor()
	// preference.DefaultFreeWeekends holds the default value on creation for the free_weekends field.
	preference.DefaultFreeWeekends = preferenceDescFreeWeekends.Default.(bool)
	// preferenceDescWeeklyFrequency is the schema descriptor for weekly_frequency field.
	preferenceDescWeeklyFrequency := preferenceFields[1].Descriptor()
	// preference.WeeklyFrequencyValidator is a validator for the "weekly_frequency" field. It is called by the builders before save.
	preference.WeeklyFrequencyValidator = preferenceDescWeeklyFrequency.Validators[0].(func(int) error)
	skillFields := schema.Skill{}.Fields()
	_ = skillFields
	// skillDescProgress is the schema descriptor for progress field.
	skillDescProgress := skillFields[2].Descriptor()
	// skill.DefaultProgress holds the default value on creation for the progress field.
	skill.DefaultProgress = skillDescProgress.Default.(int)
	// skillDescDuration is the schema descriptor for duration field.
	skillDescDuration := skillFields[3].Descriptor()
	// skill.DefaultDuration holds the default value on creation for the duration field.
	skill.DefaultDuration = skillDescDuration.Default.(int)
	// skillDescCreatedAt is the schema descriptor for created_at field.
	skillDescCreatedAt := skillFields[4].Descriptor()
	// skill.DefaultCreatedAt holds the default value on creation for the created_at field.
	skill.DefaultCreatedAt = skillDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUUID is the schema descriptor for uuid field.
	userDescUUID := userFields[4].Descriptor()
	// user.DefaultUUID holds the default value on creation for the uuid field.
	user.DefaultUUID = userDescUUID.Default.(func() uuid.UUID)
	// userDescPremium is the schema descriptor for premium field.
	userDescPremium := userFields[5].Descriptor()
	// user.DefaultPremium holds the default value on creation for the premium field.
	user.DefaultPremium = userDescPremium.Default.(bool)
}
