package service

import (
	"scripbox/hackathon/errors"
	"scripbox/hackathon/lib/database"
	"scripbox/hackathon/model"
)

//UserService structure
type UserService struct {
	DbClient database.DBClientInterface
}

//GetUserDetails to get user details
func (us *UserService) GetUserDetails(userID int) (model.User, *errors.ServiceError) {
	user := model.User{}
	return user, nil
}

type UserServiceInyerface interface {
	GetUserDetails(userID int) (model.User, *errors.ServiceError)
}
