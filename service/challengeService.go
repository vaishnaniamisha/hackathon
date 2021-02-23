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

const (
	internalServerErr = "Internal Server Error"
)

//ChallengeServiceInterface interface
type ChallengeServiceInterface interface {
	GetChallengeTagList() ([]string, *errors.ServiceError)
	ValidateTag(tag string) *errors.ServiceError
	AddChallenge(challenge model.Challenge) *errors.ServiceError
}

//GetChallengeTagList to get list of tags
func (cs ChallengeService) GetChallengeTagList() ([]string, *errors.ServiceError) {
	tags := []string{}
	tagList, err := cs.DbClient.TagList()
	if err != nil {
		log.Println("Error Getting TagList :", err.Error())
		return tags, &errors.ServiceError{
			Code:         http.StatusInternalServerError,
			ErrorMessage: internalServerErr,
		}
	}
	for _, tag := range tagList {
		tags = append(tags, tag.Tag)
	}
	return tags, nil
}

//AddChallenge to
func (cs ChallengeService) AddChallenge(challenge model.Challenge) *errors.ServiceError {
	err := cs.DbClient.CreateChallenge(challenge)
	if err != nil {
		log.Println("Error while creating challenge", err.Error())
		return &errors.ServiceError{
			Code:         http.StatusInternalServerError,
			ErrorMessage: internalServerErr,
		}
	}
	return nil
}

//ValidateTag to verify if tag exist
func (cs ChallengeService) ValidateTag(tag string) *errors.ServiceError {
	tagExist := cs.DbClient.TagExist(tag)
	if !tagExist {
		return &errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: Please provide valid tag",
		}
	}
	return nil
}
