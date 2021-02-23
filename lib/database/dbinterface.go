package database

import "scripbox/hackathon/model"

//DBClientInterface interface to include database interection method
type DBClientInterface interface {
	DBConnect() error
	TagList() ([]model.Tags, error)
	CreateChallenge(model.Challenge) error
	TagExist(string) bool
	GetUserDetails(userID int) (model.User, error)
	GetChallengeDetails(challengeID int) (model.Challenge, error)
	UpdateChallenge(challenge model.Challenge) (model.Challenge, error)
	GetAllChallenges(params map[string][]string) ([]model.Challenge, error)
}
