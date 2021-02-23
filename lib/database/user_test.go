package database_test

import (
	"regexp"
	"scripbox/hackathon/lib/database"
	"scripbox/hackathon/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserDetails(t *testing.T) {
	db, mock := SetupSqlTestDb(t)
	defer db.Close()
	client := &database.DBClient{}
	client.GormDB = db
	userID := 1001
	user := model.User{
		ID:   1001,
		Name: "User 1",
	}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"  WHERE ("ID" = $1)`)).WithArgs(userID).WillReturnRows(sqlmock.NewRows([]string{"ID", "Name"}).AddRow(user.ID, user.Name))
	res, err := client.GetUserDetails(userID)
	assert.Nil(t, err)
	assert.Equal(t, user, res)
}
