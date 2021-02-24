package database

import (
	"fmt"
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
	sql := `select ch."ID" As "ChallengeID" ,ch."Title",ch."Description",ch."Tag",ch."VoteCount",ch."CreatedBy",ch."CreatedDate",ch."IsDeleted","users"."Name" as "UserName","users"."ID" As "UserID"
	from "Challenges" ch LEFT JOIN "ChallengeCollabration" ON ch."ID" = "ChallengeCollabration"."ChallengeId"
		LEFT JOIN "users" ON "ChallengeCollabration"."UserId" = "users"."ID"`
	query := dc.GormDB.Debug()
	challenges := []model.Challenge{}
	type challengeDb struct {
		ChallengeID int       `gorm:"column:ChallengeID"`
		Title       string    `gorm:"column:Title"`
		Description string    `gorm:"column:Description"`
		Tag         string    `gorm:"column:Tag"`
		VoteCount   int       `gorm:"column:VoteCount"`
		CreatedBy   int       `gorm:"column:CreatedBy"`
		CreatedDate time.Time `gorm:"column:CreatedDate"`
		IsDeleted   bool      `gorm:"column:IsDeleted"`
		UserID      int       `gorm:"column:UserID"`
		UserName    string    `gorm:"column:UserName"`
	}
	challengesDb := []challengeDb{}

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
		if sortBy != "" {
			sql = sql + " ORDER BY " + sortBy
		}
		// query = query.Order(sortBy)
	}
	err := query.Raw(sql).Scan(&challengesDb).Error
	fmt.Println("challengesDb", len(challengesDb))
	fmt.Println("challengesDb", challengesDb)

	challengMap := make(map[int]model.Challenge)
	for _, c := range challengesDb {

		if obj, found := challengMap[c.ChallengeID]; found {
			if c.UserID != 0 {
				user := model.User{}
				user.ID = c.UserID
				user.Name = c.UserName
				obj.Collabrators = append(obj.Collabrators, user)
			}
			challengMap[c.ChallengeID] = obj
		} else {
			challenge := model.Challenge{}
			challenge.ID = c.ChallengeID
			challenge.Title = c.Title
			challenge.Description = c.Description
			challenge.Tag = c.Tag
			challenge.CreatedBy = c.CreatedBy
			challenge.CreatedDate = c.CreatedDate
			challenge.VoteCount = c.VoteCount
			challenge.IsDeleted = c.IsDeleted
			if c.UserID != 0 {
				user := model.User{}
				user.ID = c.UserID
				user.Name = c.UserName
				challenge.Collabrators = append(challenge.Collabrators, user)
			}
			challengMap[c.ChallengeID] = challenge
		}

	}
	fmt.Println("Map -- ", challengMap)
	for _, challenge := range challengMap {
		challenges = append(challenges, challenge)
	}
	return challenges, err
}

//CreateChallengeCollabration to add collabration
func (dc *DBClient) CreateChallengeCollabration(collabration model.ChallengeCollabration) error {
	err := dc.GormDB.Debug().Where(`"UserId" = ? and "ChallengeId" = ?`, collabration.UserID, collabration.ChallengeID).FirstOrCreate(&collabration).Error
	return err
}
