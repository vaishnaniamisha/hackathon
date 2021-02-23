package service_test

import (
	"errors"
	"scripbox/hackathon/lib/database"
	"scripbox/hackathon/model"
	"scripbox/hackathon/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserDetailsSuccess(t *testing.T) {
	userID := 1001
	user := model.User{
		ID:   1001,
		Name: "User 1",
	}
	mockDBRepo := new(database.MockRepo)
	userService := service.UserService{
		DbClient: mockDBRepo,
	}

	mockDBRepo.On("GetUserDetails", userID).Return(user, nil)
	res, err := userService.GetUserDetails(userID)
	assert.Nil(t, err)
	assert.Equal(t, user, res)

}
func TestGetUserDetailsError(t *testing.T) {
	userID := 1001

	mockDBRepo := new(database.MockRepo)
	userService := service.UserService{
		DbClient: mockDBRepo,
	}

	mockDBRepo.On("GetUserDetails", userID).Return(model.User{}, errors.New("db error"))
	_, err := userService.GetUserDetails(userID)
	assert.NotNil(t, err)

}
