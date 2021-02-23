package database

import (
	"scripbox/hackathon/model"

	"github.com/stretchr/testify/mock"
)

//MockRepo for UT
type MockRepo struct {
	mock.Mock
}

//TagList mock
func (m *MockRepo) TagList() ([]model.Tags, error) {
	args := m.Mock.Called()
	return args.Get(0).([]model.Tags), args.Error(1)
}

//DBConnect mock
func (m *MockRepo) DBConnect() error {
	args := m.Mock.Called()
	return args.Error(0)
}

//CreateChallenge mock
func (m *MockRepo) CreateChallenge(challenge model.Challenge) error {
	args := m.Mock.Called(challenge)
	return args.Error(0)
}

//TagExist mock
func (m *MockRepo) TagExist(tag string) bool {
	args := m.Mock.Called(tag)
	return args.Bool(0)
}

//GetUserDetails mock
func (m *MockRepo) GetUserDetails(userID int) (model.User, error) {
	args := m.Mock.Called(userID)
	return args.Get(0).(model.User), args.Error(1)
}

//GetChallengeDetails mock
func (m *MockRepo) GetChallengeDetails(challengeID int) (model.Challenge, error) {
	args := m.Mock.Called(challengeID)
	return args.Get(0).(model.Challenge), args.Error(1)
}

//UpdateChallenge mock
func (m *MockRepo) UpdateChallenge(challenge model.Challenge) (model.Challenge, error) {
	args := m.Mock.Called(challenge)
	return args.Get(0).(model.Challenge), args.Error(1)
}
