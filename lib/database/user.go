package database

import "scripbox/hackathon/model"

//UserDetails to get user deatails from DB
func (dc *DBClient) UserDetails(userID int) (model.User, error) {
	return model.User{}, nil
}
