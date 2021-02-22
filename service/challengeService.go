package service

import (
	"scripbox/hackathon/errors"
	"scripbox/hackathon/lib/database"
)

//ChallengeService strucure
type ChallengeService struct {
	DbClient database.DBClientInterface
}

//ChallengeServiceInterface interface
type ChallengeServiceInterface interface {
	GetChallengeTagList() ([]string, *errors.ServiceError)
}

//GetChallengeTagList to get list of tags
func (cs ChallengeService) GetChallengeTagList() ([]string, *errors.ServiceError) {
	return []string{}, nil
}
