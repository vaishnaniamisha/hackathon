package database_test

import (
	"database/sql/driver"
	"regexp"
	"scripbox/hackathon/lib/database"
	"scripbox/hackathon/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

var createdTime = AnyTime{}

func SetupSqlTestDb(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	gormdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	return gormdb, mock
}
func TestTagList(t *testing.T) {
	db, mock := SetupSqlTestDb(t)
	defer db.Close()
	client := &database.DBClient{}
	client.GormDB = db
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "Tags"`)).WillReturnRows(sqlmock.NewRows([]string{"Tag"}).AddRow("test"))
	tags, err := client.TagList()
	assert.Nil(t, err)
	assert.Equal(t, []model.Tags{{Tag: "test"}}, tags)
}

func TestTagExist(t *testing.T) {
	db, mock := SetupSqlTestDb(t)
	defer db.Close()
	client := &database.DBClient{}
	client.GormDB = db
	mock.ExpectExec(regexp.QuoteMeta(`SELECT * FROM "Tags" where "Tag" = $1;`)).WithArgs("tag").WillReturnResult(sqlmock.NewResult(1, 1))
	isExist := client.TagExist("tag")
	assert.True(t, isExist)
}
func TestCreateChallenge(t *testing.T) {
	db, mock := SetupSqlTestDb(t)
	defer db.Close()
	client := &database.DBClient{}
	client.GormDB = db
	challenge := model.Challenge{
		Title:       "Challenge 1",
		Description: "Description 1",
		Tag:         "tag1",
		CreatedBy:   1001,
	}
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "Challenges" ("Title","Description","Tag","VoteCount","CreatedBy","CreatedDate","IsDeleted") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "Challenges"."ID"`)).
		WithArgs(challenge.Title, challenge.Description, challenge.Tag, challenge.VoteCount, challenge.CreatedBy, createdTime, false).
		WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow(1001))
	mock.ExpectCommit()
	err := client.CreateChallenge(challenge)
	assert.Nil(t, err)
}
