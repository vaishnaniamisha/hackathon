package service

import (
	"log"
	"net/http"
	"scripbox/hackathon/errors"
	"scripbox/hackathon/lib/database"
	"scripbox/hackathon/model"
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
	tags := []string{}
	tagList, err := cs.DbClient.TagList()
	if err != nil {
		log.Println("Error Getting TagList :", err.Error())
		return tags, &errors.ServiceError{
			Code:         http.StatusInternalServerError,
			ErrorMessage: "Internal Server Error",
		}
	}
	for _, tag := range tagList {
		tags = append(tags, tag.Tag)
	}
	return tags, nil
}

//AddChallenge to
func (cs ChallengeService) AddChallenge(challenge model.Challenge) *errors.ServiceError {
	return nil
}
