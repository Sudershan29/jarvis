package models

import (
	"backend/ent"
	"backend/ent/proposal"
	"backend/ent/task"
	"backend/ent/user"
	"backend/src/helpers"
	"backend/src/lib"
	"encoding/json"
	"math"
	"strings"
	"time"
)

type TaskModel struct {
	Task *ent.Task
}

type TaskJSON struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Duration       int       `json:"duration"`
	Deadline       time.Time `json:"deadline"`
	Categories     []string  `json:"categories"`
	TimePreference []string  `json:"timepreference"`
}

func (s TaskModel) Categories() ([]*ent.Category, error) {
	dbClient := lib.DbCtx
	return s.Task.QueryCategories().All(dbClient.Context)
}

func (s TaskModel) TimePreferences() ([]*ent.TimePreference, error) {
	dbClient := lib.DbCtx
	return s.Task.QueryTimePreferences().All(dbClient.Context)
}

func (s TaskModel) Marshal() TaskJSON {
	catArr := make([]string, 0)
	categories, _ := s.Categories()
	for _, cat := range categories {
		catArr = append(catArr, cat.Name)
	}
	timePrefArr := make([]string, 0)
	timePreferences, _ := s.TimePreferences()
	for _, tp := range timePreferences {
		timePrefArr = append(timePrefArr, tp.Day)
	}

	return TaskJSON{Id: s.Task.ID, Name: s.Task.Name, Description: s.Task.Description, Duration: s.Task.Duration, Deadline: s.Task.Deadline, Categories: catArr, TimePreference: timePrefArr}
}

/* * * * * * * * * * * *

		APIs

* * * * * * * * * * * * */

func TaskCreate(name, description string, duration int, deadline time.Time, categories []string, timePreferences []string, currUser *JwtUser) (*TaskModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	sOrm := dbClient.Client.Task.
		Create().
		SetName(name).
		SetDuration(duration).
		SetDeadline(deadline).
		SetUser(currUser.Model.User).
		SetDescription(description)

	for _, cat := range categories {
		catModel, err := CategoryFindOrCreate(cat, currUser)
		if err != nil {
			continue
		}
		sOrm.AddCategories(catModel.Category)
	}

	for _, timePref := range timePreferences {
		timePrefModel, err := TimePreferenceFindOrCreate(timePref)
		if err != nil {
			continue
		}
		sOrm.AddTimePreferences(timePrefModel.TimePreference)
	}

	s, err := sOrm.Save(dbClient.Context)

	if err != nil {
		return nil, err
	}

	task := TaskModel{s}
	return &task, nil
}

func TaskShowAll(currUser *JwtUser) ([]*TaskModel, error) {
	dbClient := lib.DbCtx
	tasks, err := dbClient.Client.Task.
		Query().
		Where(task.HasUserWith(user.UUID(currUser.UserId))).
		Order(ent.Desc("created_at")).
		All(dbClient.Context)

	if err != nil {
		return make([]*TaskModel, 0), err
	}

	result := make([]*TaskModel, 0)
	for _, task := range tasks {
		result = append(result, &TaskModel{task})
	}
	return result, nil
}

func TaskDelete(taskID int, currUser *JwtUser) error {
	dbClient := lib.DbCtx
	// Find the task by ID and ensure it belongs to the current user
	task, err := dbClient.Client.Task.
		Query().
		Where(task.ID(taskID), task.HasUserWith(user.UUID(currUser.UserId))).
		Only(dbClient.Context)
	if err != nil {
		return err
	}

	// Delete the task
	return dbClient.Client.Task.DeleteOne(task).Exec(dbClient.Context)
}

func TaskFind(taskID int, currUser *JwtUser) (*TaskModel, error) {
	dbClient := lib.DbCtx
	task, err := dbClient.Client.Task.
		Query().
		Where(task.ID(taskID), task.HasUserWith(user.UUID(currUser.UserId))).
		Only(dbClient.Context)

	if err != nil {
		return nil, err
	}

	return &TaskModel{task}, nil
}

func (t *TaskModel) CancelProposal(proposalID int) error {
	proposal, err := ProposalFindByTaskIDAndProposalID(t.Task.ID, proposalID)

	if err == nil {
		return proposal.Cancel()
	}

	return err
}

func (t *TaskModel) ProposalsWithTimeFilter(startTime time.Time, endTime time.Time) ([]*ProposalModel, error) {
	dbClient := lib.DbCtx
	ps, err := dbClient.Client.Proposal.
		Query().
		Where(
			proposal.HasTaskWith(task.ID(t.Task.ID)),
			proposal.ScheduledForGTE(startTime),
			proposal.ScheduledForLTE(endTime),
			proposal.StatusNEQ(proposal.StatusDeleted),
		).
		All(dbClient.Context)

	if err != nil {
		return nil, err
	}

	proposals := make([]*ProposalModel, 0)

	for _, p := range ps {
		proposals = append(proposals, &ProposalModel{p})
	}

	return proposals, nil
}

func (t *TaskModel) HoursLeft(currDate time.Time) int {
	thisWeekProposals, _ := t.ProposalsWithTimeFilter(t.Task.CreatedAt, helpers.EndOfWeek(currDate))

	scheduled := 0
	for _, p := range thisWeekProposals {
		if p.Proposal.Status == proposal.StatusPending {
			scheduled += p.Proposal.AllocatedDuration
		} else if p.Proposal.Status == proposal.StatusDone {
			scheduled += p.Proposal.AchievedDuration
		}
	}

	return t.Task.Duration - scheduled
}

func (t *TaskModel) GetTaskTimePreferences() ([]string, error) {
	dbClient := lib.DbCtx
	timePrefs, err := dbClient.Client.Task.
		Query().
		Where(task.ID(t.Task.ID)).
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

func (t *TaskModel) PreferredDaysLeft(currDate time.Time) int {
	preferredDays, _ := t.GetTaskTimePreferences()
	if len(preferredDays) == 0 {
		preferredDays = helpers.WEEK
	}
	scheduledAlready, _ := t.ProposalsWithTimeFilter(t.Task.CreatedAt, helpers.EndOfWeek(currDate))
	idealNumberOfSchedules := t.idealNumberOfSchedules() - len(scheduledAlready)
	if !t.Task.Deadline.IsZero() {
		idealNumberOfSchedules = int(math.Min(float64(helpers.NumberOfPreferredDaysBetween(currDate, t.Task.Deadline, preferredDays)), float64(idealNumberOfSchedules)))
	}
	return idealNumberOfSchedules
}

func (t *TaskModel) idealNumberOfSchedules() int {
	// Assuming Workload to be split into 3 hr blocks
	return t.Task.Duration / 3
}

func (t *TaskModel) PreferredMaxHours() int {
	return 6
}

func (t *TaskModel) PreferredMinHours() int {
	return 2
}

func (t *TaskModel) IsPreferredDay(day string) bool {
	days, _ := t.GetTaskTimePreferences()
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

func (t *TaskModel) String() string {
	data, _ := json.Marshal(t.Marshal())
	return string(data)
}

func (t *TaskModel) AddProposal(proposal *ProposalModel) bool {
	dbClient := lib.DbCtx
	_, err := dbClient.Client.Task.UpdateOne(t.Task).
		AddProposals(proposal.Proposal).
		Save(dbClient.Context)

	return err != nil
}

func (t *TaskModel) Name() string {
	return t.Task.Name
}

func (t *TaskModel) UserUUID() string {
	return t.Task.QueryUser().OnlyX(lib.DbCtx.Context).UUID.String()
}

func (t *TaskModel) HoursNeeded(currDate time.Time) float64 {
	hoursLeft := t.HoursLeft(currDate)
	daysLeft := t.PreferredDaysLeft(currDate)

	if daysLeft <= 0 {
		return 0 // check how long is remaining ig?
	}

	avg := hoursLeft / daysLeft

	if t.IsPreferredDay(currDate.Weekday().String()) {
		setMinimum := math.Max(float64(t.PreferredMinHours()), float64(avg))
		capMaxmium := math.Min(setMinimum, float64(t.PreferredMaxHours()))
		return capMaxmium
	} else {
		return 0
	}
}
