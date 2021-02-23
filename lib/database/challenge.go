package database

import (
	"scripbox/hackathon/model"
	"time"
)

//TagList to get all tags
func (dc *DBClient) TagList() ([]model.Tags, error) {
	tags := []model.Tags{}
	err := dc.GormDB.Find(&tags).Error
	return tags, err
}

//CreateChallenge to create new challenge in db
func (dc *DBClient) CreateChallenge(challenge model.Challenge) error {
	challenge.CreatedDate = time.Now()
	err := dc.GormDB.Create(&challenge).Error
	return err
}

//TagExist to check if tag exist
func (dc *DBClient) TagExist(tag string) bool {
	query := `SELECT * FROM "Tags" where "Tag" = ?;`
	rows := dc.GormDB.Exec(query, tag).RowsAffected
	return rows > 0
}

//GetChallengeDetails query
func (dc *DBClient) GetChallengeDetails(challengeID int) (model.Challenge, error) {
	challenge := model.Challenge{}
	err := dc.GormDB.Debug().Where(`"ID" = ?`, challengeID).Find(&challenge).Error
	return challenge, err
}

//UpdateChallenge to update challenge details
func (dc *DBClient) UpdateChallenge(challenge model.Challenge) (model.Challenge, error) {
	err := dc.GormDB.Debug().Save(&challenge).Error
	return challenge, err
}
