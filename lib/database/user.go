package database

import (
	"scripbox/hackathon/model"
)

//GetUserDetails to get user deatails from DB
func (dc *DBClient) GetUserDetails(userID int) (model.User, error) {
	user := model.User{}
	err := dc.GormDB.Debug().Where(`"ID" = ?`, userID).Find(&user).Error
	return user, err
}
