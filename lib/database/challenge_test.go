package database_test

import (
	"regexp"
	"scripbox/hackathon/lib/database"
	"scripbox/hackathon/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

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
