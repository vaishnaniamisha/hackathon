package database

import "scripbox/hackathon/model"

//GetUserDetails to get user deatails from DB
func (dc *DBClient) GetUserDetails(userID int) (model.User, error) {
	return model.User{}, nil
}
