package service

import (
	"log"
	"net/http"
	"scripbox/hackathon/errors"
	"scripbox/hackathon/lib/database"
	"scripbox/hackathon/model"

	"github.com/jinzhu/gorm"
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
	GetChallengeDetails(challengeID int) (model.Challenge, *errors.ServiceError)
	UpvoteChallenge(challenge model.Challenge) (model.Challenge, *errors.ServiceError)
	ListAllChallenges(params map[string][]string) ([]model.Challenge, *errors.ServiceError)
	UpdateCollabration(userID int, challengeID int) *errors.ServiceError
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

//GetChallengeDetails to get challenge
func (cs ChallengeService) GetChallengeDetails(challengeID int) (model.Challenge, *errors.ServiceError) {
	challenge, err := cs.DbClient.GetChallengeDetails(challengeID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return challenge, &errors.ServiceError{
				Code:         http.StatusNotFound,
				ErrorMessage: "Challenge does not exist",
			}
		}
		log.Println("Error getting Challenge:", err.Error())
		return challenge, &errors.ServiceError{
			Code:         http.StatusInternalServerError,
			ErrorMessage: internalServerErr,
		}
	}
	return challenge, nil
}

//UpvoteChallenge to upvote challenge
func (cs ChallengeService) UpvoteChallenge(challenge model.Challenge) (model.Challenge, *errors.ServiceError) {
	challenge.VoteCount = challenge.VoteCount + 1
	challenge, err := cs.DbClient.UpdateChallenge(challenge)
	if err != nil {
		log.Println("Error updating Votes:", err.Error())
		return challenge, &errors.ServiceError{
			Code:         http.StatusInternalServerError,
			ErrorMessage: internalServerErr,
		}
	}
	return challenge, nil
}

//ListAllChallenges to get all challenges
func (cs ChallengeService) ListAllChallenges(params map[string][]string) ([]model.Challenge, *errors.ServiceError) {
	challenges, err := cs.DbClient.GetAllChallenges(params)
	if err != nil {
		log.Println("Error getting all challenges:", err.Error())
		return challenges, &errors.ServiceError{
			Code:         http.StatusInternalServerError,
			ErrorMessage: internalServerErr,
		}
	}
	return challenges, nil
}

//UpdateCollabration to add collabration
func (cs ChallengeService) UpdateCollabration(userID int, challengeID int) *errors.ServiceError {
	collabration := model.ChallengeCollabration{
		UserID:      userID,
		ChallengeID: challengeID,
	}
	err := cs.DbClient.CreateChallengeCollabration(collabration)
	if err != nil {
		log.Println("Error creating collabration:", err.Error())
		return &errors.ServiceError{
			Code:         http.StatusInternalServerError,
			ErrorMessage: internalServerErr,
		}
	}
	return nil
}
