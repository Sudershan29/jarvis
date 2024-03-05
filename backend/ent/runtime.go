// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/goal"
	"backend/ent/hobby"
	"backend/ent/meeting"
	"backend/ent/preference"
	"backend/ent/proposal"
	"backend/ent/schema"
	"backend/ent/skill"
	"backend/ent/task"
	"backend/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	goalFields := schema.Goal{}.Fields()
	_ = goalFields
	// goalDescCreatedAt is the schema descriptor for created_at field.
	goalDescCreatedAt := goalFields[2].Descriptor()
	// goal.DefaultCreatedAt holds the default value on creation for the created_at field.
	goal.DefaultCreatedAt = goalDescCreatedAt.Default.(func() time.Time)
	hobbyFields := schema.Hobby{}.Fields()
	_ = hobbyFields
	// hobbyDescCreatedAt is the schema descriptor for created_at field.
	hobbyDescCreatedAt := hobbyFields[2].Descriptor()
	// hobby.DefaultCreatedAt holds the default value on creation for the created_at field.
	hobby.DefaultCreatedAt = hobbyDescCreatedAt.Default.(func() time.Time)
	meetingFields := schema.Meeting{}.Fields()
	_ = meetingFields
	// meetingDescDuration is the schema descriptor for duration field.
	meetingDescDuration := meetingFields[4].Descriptor()
	// meeting.DefaultDuration holds the default value on creation for the duration field.
	meeting.DefaultDuration = meetingDescDuration.Default.(int)
	// meetingDescCreatedAt is the schema descriptor for created_at field.
	meetingDescCreatedAt := meetingFields[5].Descriptor()
	// meeting.DefaultCreatedAt holds the default value on creation for the created_at field.
	meeting.DefaultCreatedAt = meetingDescCreatedAt.Default.(func() time.Time)
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
	proposalFields := schema.Proposal{}.Fields()
	_ = proposalFields
	// proposalDescAllocatedDuration is the schema descriptor for allocated_duration field.
	proposalDescAllocatedDuration := proposalFields[1].Descriptor()
	// proposal.AllocatedDurationValidator is a validator for the "allocated_duration" field. It is called by the builders before save.
	proposal.AllocatedDurationValidator = proposalDescAllocatedDuration.Validators[0].(func(int) error)
	// proposalDescAchievedDuration is the schema descriptor for achieved_duration field.
	proposalDescAchievedDuration := proposalFields[2].Descriptor()
	// proposal.AchievedDurationValidator is a validator for the "achieved_duration" field. It is called by the builders before save.
	proposal.AchievedDurationValidator = proposalDescAchievedDuration.Validators[0].(func(int) error)
	// proposalDescCreatedAt is the schema descriptor for created_at field.
	proposalDescCreatedAt := proposalFields[5].Descriptor()
	// proposal.DefaultCreatedAt holds the default value on creation for the created_at field.
	proposal.DefaultCreatedAt = proposalDescCreatedAt.Default.(func() time.Time)
	// proposalDescUpdatedAt is the schema descriptor for updated_at field.
	proposalDescUpdatedAt := proposalFields[6].Descriptor()
	// proposal.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	proposal.DefaultUpdatedAt = proposalDescUpdatedAt.Default.(func() time.Time)
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
	// skillDescDurationAchieved is the schema descriptor for duration_achieved field.
	skillDescDurationAchieved := skillFields[4].Descriptor()
	// skill.DefaultDurationAchieved holds the default value on creation for the duration_achieved field.
	skill.DefaultDurationAchieved = skillDescDurationAchieved.Default.(int)
	// skillDescCreatedAt is the schema descriptor for created_at field.
	skillDescCreatedAt := skillFields[5].Descriptor()
	// skill.DefaultCreatedAt holds the default value on creation for the created_at field.
	skill.DefaultCreatedAt = skillDescCreatedAt.Default.(func() time.Time)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescDuration is the schema descriptor for duration field.
	taskDescDuration := taskFields[2].Descriptor()
	// task.DefaultDuration holds the default value on creation for the duration field.
	task.DefaultDuration = taskDescDuration.Default.(int)
	// taskDescDurationAchieved is the schema descriptor for duration_achieved field.
	taskDescDurationAchieved := taskFields[3].Descriptor()
	// task.DefaultDurationAchieved holds the default value on creation for the duration_achieved field.
	task.DefaultDurationAchieved = taskDescDurationAchieved.Default.(int)
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[4].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
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
