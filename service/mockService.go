package service

import (
	"scripbox/hackathon/errors"
	"scripbox/hackathon/model"

	"github.com/stretchr/testify/mock"
)

//MockChallengeService mock struct
type MockChallengeService struct {
	mock.Mock
}

//MockUserService mock struct
type MockUserService struct {
	mock.Mock
}

//GetChallengeTagList mock
func (m MockChallengeService) GetChallengeTagList() ([]string, *errors.ServiceError) {
	args := m.Mock.Called()
	return args.Get(0).([]string), args.Get(1).(*errors.ServiceError)
}

//AddChallenge mock
func (m MockChallengeService) AddChallenge(challenge model.Challenge) *errors.ServiceError {
	args := m.Mock.Called(challenge)
	return args.Get(0).(*errors.ServiceError)
}

//ValidateTag mock
func (m MockChallengeService) ValidateTag(tag string) *errors.ServiceError {
	args := m.Mock.Called(tag)
	return args.Get(0).(*errors.ServiceError)
}

//GetUserDetails mock
func (m MockUserService) GetUserDetails(userID int) (model.User, *errors.ServiceError) {
	args := m.Mock.Called(userID)
	return args.Get(0).(model.User), args.Get(1).(*errors.ServiceError)
}
