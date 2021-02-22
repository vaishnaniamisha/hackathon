package database

import "scripbox/hackathon/model"

//TagList to get all tags
func (dc *DBClient) TagList() ([]model.Tags, error) {
	tags := []model.Tags{}
	err := dc.GormDB.Debug().Find(&tags).Error
	return tags, err
}
