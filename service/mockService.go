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
