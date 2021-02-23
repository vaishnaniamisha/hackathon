package service

import (
	"log"
	"net/http"
	"scripbox/hackathon/errors"
	"scripbox/hackathon/lib/database"
	"scripbox/hackathon/model"

	"github.com/jinzhu/gorm"
)

//UserService structure
type UserService struct {
	DbClient database.DBClientInterface
}

//UserServiceInyerface interface
type UserServiceInyerface interface {
	GetUserDetails(userID int) (model.User, *errors.ServiceError)
}

//GetUserDetails to get user details
func (us *UserService) GetUserDetails(userID int) (model.User, *errors.ServiceError) {
	user, err := us.DbClient.GetUserDetails(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, &errors.ServiceError{
				Code:         http.StatusNotFound,
				ErrorMessage: "UserId does not exist",
			}
		}
		log.Println("Error getting user details : ", err.Error())
		return model.User{}, &errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: internalServerErr,
		}
	}

	return user, nil
}
