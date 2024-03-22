package models

import (
	"backend/ent"
	"backend/ent/proposal"
	"backend/ent/skill"
	"backend/ent/user"
	"backend/src/helpers"
	"backend/src/lib"
	"encoding/json"
	"math"
	"strings"
	"time"
)

type SkillModel struct {
	Skill *ent.Skill
}

type SkillJSON struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Level          string   `json:"level"`
	Duration       int      `json:"duration"`
	Categories     []string `json:"categories"`
	TimePreference []string `json:"timepreference"`
}

func (s SkillModel) Categories() ([]*ent.Category, error) {
	dbClient := lib.DbCtx
	return s.Skill.QueryCategories().All(dbClient.Context)
}

func (s SkillModel) TimePreference() ([]*ent.TimePreference, error) {
	dbClient := lib.DbCtx
	return s.Skill.QueryTimePreferences().All(dbClient.Context)
}

func (s SkillModel) Marshal() SkillJSON {
	catArr := make([]string, 0)
	categories, _ := s.Categories()
	for _, cat := range categories {
		catArr = append(catArr, cat.Name)
	}
	timePrefArr := make([]string, 0)
	timePreferences, _ := s.TimePreference()
	for _, tp := range timePreferences {
		timePrefArr = append(timePrefArr, tp.Day)
	}
	return SkillJSON{Id: s.Skill.ID, Name: s.Skill.Name, Level: s.Skill.Level, Duration: s.Skill.Duration, Categories: catArr, TimePreference: timePrefArr}
}

/* * * * * * * * * * * *

		APIs

* * * * * * * * * * * * */

func SkillCreate(name, level string, duration int, categories []string, timePreference []string, currUser *JwtUser) (*SkillModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	sOrm := dbClient.Client.Skill.
		Create().
		SetName(name).
		SetUser(currUser.Model.User).
		SetDuration(duration).
		SetLevel(level)

	for _, cat := range categories {
		catModel, err := CategoryFindOrCreate(cat, currUser)
		if err != nil {
			continue
		}
		sOrm.AddCategories(catModel.Category)
	}

	for _, tp := range timePreference {
		tpModel, err := TimePreferenceFindOrCreate(tp)
		if err != nil {
			continue
		}
		sOrm.AddTimePreferences(tpModel.TimePreference)
	}

	s, err := sOrm.Save(dbClient.Context)

	if err != nil {
		return nil, err
	}

	skill := SkillModel{s}
	return &skill, nil
}

func SkillShowAll(currUser *JwtUser) ([]*SkillModel, error) {
	dbClient := lib.DbCtx
	skills, err := dbClient.Client.Skill.
		Query().
		Where(skill.HasUserWith(user.UUID(currUser.UserId))).
		All(dbClient.Context)

	if err != nil {
		return make([]*SkillModel, 0), err
	}

	result := make([]*SkillModel, 0)
	for _, skill := range skills {
		result = append(result, &SkillModel{skill})
	}
	return result, nil
}

func SkillDelete(skillID int, currUser *JwtUser) error {
	dbClient := lib.DbCtx
	err := dbClient.Client.Skill.
		DeleteOneID(skillID).
		Where(skill.HasUserWith(user.UUID(currUser.UserId))).
		Exec(dbClient.Context)

	if err != nil {
		return err
	}

	return nil
}

func SkillFind(skillID int, currUser *JwtUser) (*SkillModel, error) {
	dbClient := lib.DbCtx
	skill, err := dbClient.Client.Skill.
		Query().
		Where(skill.ID(skillID), skill.HasUserWith(user.UUID(currUser.UserId))).
		Only(dbClient.Context)

	if err != nil {
		return nil, err
	}

	return &SkillModel{skill}, nil
}

func (s *SkillModel) CancelProposal(proposalID int) error {
	proposal, err := ProposalFindBySkillIDAndProposalID(s.Skill.ID, proposalID)

	if err == nil {
		return proposal.Cancel()
	}

	return err
}

func (s *SkillModel) ProposalsWithTimeFilter(startTime time.Time, endTime time.Time) ([]*ProposalModel, error) {
	dbClient := lib.DbCtx
	ps, err := dbClient.Client.Proposal.
		Query().
		Where(
			proposal.HasSkillWith(skill.ID(s.Skill.ID)),
			proposal.ScheduledForGTE(startTime),
			proposal.ScheduledForLTE(endTime),
			proposal.StatusNEQ(proposal.StatusDeleted),
		).
		All(dbClient.Context)

	if err != nil {
		return nil, err
	}

	var proposals []*ProposalModel

	for _, p := range ps {
		proposals = append(proposals, &ProposalModel{p})
	}

	return proposals, nil
}

func (s *SkillModel) HoursLeft(d time.Time) int {
	thisWeekProposals, _ := s.ProposalsWithTimeFilter(helpers.StartOfWeek(d), helpers.EndOfWeek(d))

	scheduled := 0
	for _, p := range thisWeekProposals {
		if p.Proposal.Status == proposal.StatusPending {
			scheduled += p.Proposal.AllocatedDuration
		} else if p.Proposal.Status == proposal.StatusDone {
			scheduled += p.Proposal.AchievedDuration
		}
	}

	return s.Skill.Duration - scheduled
}

func (s *SkillModel) PreferredDaysLeft(d time.Time) int {
	preferredDays, _ := s.GetSkillTimePreferences()
	if len(preferredDays) == 0 {
		return len(helpers.NextDays(d))
	} else {
		return helpers.NumberOfDaysLeftThisWeek(d, preferredDays)
	}
}

func (s *SkillModel) PreferredMaxHours() int {
	return 6 // TODO: Note vary it based on user
}

func (s *SkillModel) PreferredMinHours() int {
	return 1 // TODO: Note vary it based on user
}

func (s *SkillModel) GetSkillTimePreferences() ([]string, error) {
	dbClient := lib.DbCtx
	timePrefs, err := dbClient.Client.Skill.
		Query().
		Where(skill.ID(s.Skill.ID)).
		QueryTimePreferences().
		All(dbClient.Context)

	if err != nil {
		return nil, err
	}

	var preferences []string
	for _, tp := range timePrefs {
		preferences = append(preferences, tp.Day)
	}

	return preferences, nil
}

func (s *SkillModel) IsPreferredDay(day string) bool {
	days, _ := s.GetSkillTimePreferences()
	if len(days) == 0 {
		return true
	} else {
		for _, d := range days {
			if strings.Title(strings.ToLower(d)) == day {
				return true
			}
		}
		return false
	}
}

func (s *SkillModel) String() string {
	data, _ := json.Marshal(s.Marshal())
	return string(data)
}

func (s *SkillModel) AddProposal(proposal *ProposalModel) bool {
	dbClient := lib.DbCtx
	_, err := dbClient.Client.Skill.UpdateOne(s.Skill).
		AddProposals(proposal.Proposal).
		Save(dbClient.Context)

	return err != nil
}

func (s *SkillModel) Name() string {
	return s.Skill.Name
}

func (s *SkillModel) HoursNeeded(t time.Time) float64 {
	hoursLeft := s.HoursLeft(t)
	daysLeft := s.PreferredDaysLeft(t)

	if daysLeft <= 0 {
		return 0 // Remaining hours? or I am past deadline?
	}

	avg := hoursLeft / daysLeft

	if s.IsPreferredDay(t.Weekday().String()) {
		setMinimum := math.Max(float64(s.PreferredMinHours()), float64(avg))
		capMaxmium := math.Min(setMinimum, float64(s.PreferredMaxHours()))
		return capMaxmium
	} else {
		return 0
	}
}

func (s *SkillModel) UserUUID() string {
	return s.Skill.QueryUser().OnlyX(lib.DbCtx.Context).UUID.String()
}
