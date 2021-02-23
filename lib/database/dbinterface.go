package database

import "scripbox/hackathon/model"

//DBClientInterface interface to include database interection method
type DBClientInterface interface {
	DBConnect() error
	TagList() ([]model.Tags, error)
	CreateChallenge(model.Challenge) error
	TagExist(string) bool
}
