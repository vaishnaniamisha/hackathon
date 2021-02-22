package service_test

import (
	"errors"
	"scripbox/hackathon/lib/database"
	"scripbox/hackathon/model"
	"scripbox/hackathon/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChallengeTagListSuccess(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	mockDBRepo.On("TagList").Return([]model.Tags{{Tag: "test"}}, nil)
	tags, err := challengeService.GetChallengeTagList()
	assert.Nil(t, err)
	assert.Equal(t, []string{"test"}, tags)

}
func TestGetChallengeTagListError(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	mockDBRepo.On("TagList").Return([]model.Tags{}, errors.New("DB Error"))
	_, err := challengeService.GetChallengeTagList()
	assert.NotNil(t, err)

}

func TestAddChallengeSuccess(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	challenge := model.Challenge{
		Title:       "Challenge 1",
		Description: "Description 1",
		Tag:         "tag1",
	}
	mockDBRepo.On("CreateChallenge", challenge).Return(nil)
	err := challengeService.AddChallenge(challenge)
	assert.Nil(t, err)
}
func TestAddChallengeError(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	challenge := model.Challenge{
		Title:       "Challenge 1",
		Description: "Description 1",
		Tag:         "tag1",
	}
	mockDBRepo.On("CreateChallenge", challenge).Return(errors.New("Db error"))
	err := challengeService.AddChallenge(challenge)
	assert.NotNil(t, err)
}
