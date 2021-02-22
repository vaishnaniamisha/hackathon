package database

import "scripbox/hackathon/model"

//TagList to get all tags
func (dc *DBClient) TagList() ([]model.Tags, error) {
	tags := []model.Tags{}
	err := dc.GormDB.Debug().Find(&tags).Error
	return tags, err
}

//CreateChallenge to create new challenge in db
func (dc *DBClient) CreateChallenge(challenge model.Challenge) error {
	return nil
}
