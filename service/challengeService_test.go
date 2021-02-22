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
