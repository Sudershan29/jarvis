package models

import (
	"time"
	"backend/ent"
	"backend/src/lib"
	"backend/ent/meeting"
	"backend/ent/user"
)

type MeetingModel struct {
	Meeting *ent.Meeting
}

type MeetingJSON struct {
	Name 	     string    `json:"name"`
	Description	 string    `json:"description"`
	Where    	 string    `json:"where"`
	Whom	     string    `json:"whom"`
	When		 time.Time `json:"when"`
	Duration	 int	   `json:"duration"`
}

func (s MeetingModel) Marshal() MeetingJSON {
	return MeetingJSON{s.Meeting.Name, s.Meeting.Description, s.Meeting.Where, s.Meeting.Whom,
					   s.Meeting.When, s.Meeting.Duration}
}


/* * * * * * * * * * * * 

		APIs

* * * * * * * * * * * * */

func MeetingCreate(name, description, where, whom string, duration int, when time.Time, currUser *JwtUser) (*MeetingModel, error) {
	currUser.Load()
	dbClient := lib.DbCtx
	sOrm := dbClient.Client.Meeting.
				Create().
				SetName(name).
				SetDescription(description).
				SetWhere(where).
				SetWhom(whom).
				SetDuration(duration).
				SetWhen(when)

	s, err := sOrm.Save(dbClient.Context)

	if err != nil { return nil, err }

	meeting := MeetingModel{s}
	return &meeting, nil
}

func MeetingShowAll(currUser *JwtUser) ([]*MeetingModel, error) {
	dbClient := lib.DbCtx
	meetings, err := dbClient.Client.Meeting.
					Query().
					Where(meeting.HasUserWith(user.UUID(currUser.UserId))).
					All(dbClient.Context)
	
	if err != nil {
		return make([]*MeetingModel, 0), err
	}

	result := make([]*MeetingModel, 0)
	for _, meeting := range meetings {
		result = append(result, &MeetingModel{meeting})
	}
	return result, nil
}
