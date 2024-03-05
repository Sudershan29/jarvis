package models

import (
	"backend/ent"
	"backend/ent/proposal"
	"backend/src/lib"
	"time"
)

type ProposalModel struct {
	Proposal *ent.Proposal
}

type ProposalJSON struct {
	Name              string    `json:"name"`
	AllocatedDuration int       `json:"allocatedDuration"`
	AchievedDuration  int       `json:"achievedDuration"`
	ScheduledFor      time.Time `json:"scheduledFor"`
	Status            string    `json:"status"`
}

func ProposalCreate(name string, allocatedDuration int, achievedDuration int, status string, scheduledFor time.Time) (*ProposalModel, error) {
	dbClient := lib.DbCtx
	proposal, err := dbClient.Client.Proposal.
		Create().
		SetName(name).
		SetAllocatedDuration(allocatedDuration).
		SetAchievedDuration(achievedDuration).
		SetScheduledFor(scheduledFor).
		SetStatus(proposal.Status(status)).
		Save(dbClient.Context)

	if err != nil {
		return nil, err
	}

	return &ProposalModel{Proposal: proposal}, nil
}

func ProposalFindOrCreate(name string) (*ProposalModel, error) {
	dbClient := lib.DbCtx
	proposal, err := dbClient.Client.Proposal.
		Query().
		Where(proposal.Name(name)).
		Only(dbClient.Context)

	if err != nil {
		// If not found, create a new Proposal
		proposal, err = dbClient.Client.Proposal.
			Create().
			SetName(name).
			Save(dbClient.Context)
		if err != nil {
			return nil, err
		}
	}

	return &ProposalModel{Proposal: proposal}, nil
}

func ProposalShowAll() ([]*ProposalModel, error) {
	dbClient := lib.DbCtx
	proposals, err := dbClient.Client.Proposal.
		Query().
		All(dbClient.Context)

	if err != nil {
		return make([]*ProposalModel, 0), err
	}

	result := make([]*ProposalModel, 0)
	for _, p := range proposals {
		result = append(result, &ProposalModel{Proposal: p})
	}
	return result, nil
}
