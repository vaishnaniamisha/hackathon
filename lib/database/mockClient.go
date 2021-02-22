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
	return args.Error(1)
}
