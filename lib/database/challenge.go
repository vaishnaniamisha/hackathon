package database

import (
	"scripbox/hackathon/model"
	"strings"
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

//GetAllChallenges to list all challneges
func (dc *DBClient) GetAllChallenges(params map[string][]string) ([]model.Challenge, error) {
	query := dc.GormDB.Debug()
	challenges := []model.Challenge{}
	if sortParam, ok := params["sortby"]; ok {
		var sortbyArr []string
		for _, sort := range sortParam {
			sortArr := strings.Split(sort, " ")
			orderParam := sortArr[0]
			order := "asc"
			if len(sortArr) <= 0 {
				continue
			}
			if len(sortArr) >= 2 && strings.EqualFold(sortArr[1], "descending") {
				order = "desc"
			}
			switch strings.ToLower(orderParam) {
			case "votecount":
				sortbyArr = append(sortbyArr, `"VoteCount" `+order)
			case "createddate":
				sortbyArr = append(sortbyArr, `"CreatedDate" `+order)
			}

		}
		sortBy := strings.Join(sortbyArr, ",")
		query = query.Order(sortBy)
	}
	err := query.Find(&challenges).Error
	return challenges, err
}

//CreateChallengeCollabration to add collabration
func (dc *DBClient) CreateChallengeCollabration(collabration model.ChallengeCollabration) error {
	err := dc.GormDB.Debug().Where(`"UserId" = ? and "ChallengeId" = ?`, collabration.UserID, collabration.ChallengeID).FirstOrCreate(&collabration).Error
	return err
}
