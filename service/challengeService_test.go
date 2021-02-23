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

func TestValidateTagSuccess(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	mockDBRepo.On("TagExist", "tag").Return(true)
	err := challengeService.ValidateTag("tag")
	assert.Nil(t, err)
}
func TestValidateTagError(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	mockDBRepo.On("TagExist", "tag").Return(false)
	err := challengeService.ValidateTag("tag")
	assert.NotNil(t, err)
}

func TestGetChallengeDetails(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	challenge := model.Challenge{
		ID:          1001,
		Title:       "Challenge 1",
		Description: "Description 1",
		Tag:         "tag1",
	}
	mockDBRepo.On("GetChallengeDetails", challenge.ID).Return(challenge, nil)
	res, err := challengeService.GetChallengeDetails(challenge.ID)
	assert.Nil(t, err)
	assert.Equal(t, challenge, res)
}
func TestGetChallengeDetailsError(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	challengeID := 1001
	mockDBRepo.On("GetChallengeDetails", challengeID).Return(model.Challenge{}, errors.New("db  error"))
	_, err := challengeService.GetChallengeDetails(challengeID)
	assert.NotNil(t, err)
}

func TestUpvoteChallenge(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	challenge := model.Challenge{
		ID:          1001,
		Title:       "Challenge 1",
		Description: "Description 1",
		Tag:         "tag1",
	}
	expectedResult := model.Challenge{
		ID:          1001,
		Title:       "Challenge 1",
		Description: "Description 1",
		Tag:         "tag1",
		VoteCount:   1,
	}
	mockDBRepo.On("UpdateChallenge", expectedResult).Return(expectedResult, nil)
	res, err := challengeService.UpvoteChallenge(challenge)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, res)
}
func TestUpvoteChallengeError(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	challenge := model.Challenge{
		ID:          1001,
		Title:       "Challenge 1",
		Description: "Description 1",
		Tag:         "tag1",
	}
	expectedResult := model.Challenge{
		ID:          1001,
		Title:       "Challenge 1",
		Description: "Description 1",
		Tag:         "tag1",
		VoteCount:   1,
	}
	mockDBRepo.On("UpdateChallenge", expectedResult).Return(expectedResult, errors.New("db error"))
	_, err := challengeService.UpvoteChallenge(challenge)
	assert.NotNil(t, err)

}

func TestListAllChallenges(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	params := make(map[string][]string)
	params["sortby"] = []string{"votecount asending", "createddate descending"}
	mockDBRepo.On("GetAllChallenges", params).Return([]model.Challenge{}, nil)
	_, err := challengeService.ListAllChallenges(params)
	assert.Nil(t, err)
}
func TestListAllChallengesError(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	params := make(map[string][]string)
	params["sortby"] = []string{"votecount asending", "createddate descending"}
	mockDBRepo.On("GetAllChallenges", params).Return([]model.Challenge{}, errors.New("db Err"))
	_, err := challengeService.ListAllChallenges(params)
	assert.NotNil(t, err)
}

func TestUpdateCollabration(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	userID := 1001
	challengeID := 1001
	collabration := model.ChallengeCollabration{
		UserID:      userID,
		ChallengeID: challengeID,
	}
	mockDBRepo.On("CreateChallengeCollabration", collabration).Return(nil)
	err := challengeService.UpdateCollabration(userID, challengeID)
	assert.Nil(t, err)
}
func TestUpdateCollabrationError(t *testing.T) {
	mockDBRepo := new(database.MockRepo)
	challengeService := service.ChallengeService{
		DbClient: mockDBRepo,
	}
	userID := 1001
	challengeID := 1001
	collabration := model.ChallengeCollabration{
		UserID:      userID,
		ChallengeID: challengeID,
	}
	mockDBRepo.On("CreateChallengeCollabration", collabration).Return(errors.New("db error"))
	err := challengeService.UpdateCollabration(userID, challengeID)
	assert.NotNil(t, err)
}
