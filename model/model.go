package model

import (
	"time"
)

//Tags struct to map Tags db table
type Tags struct {
	Tag string `gorm:"column:Tag"`
}

//TableName for Tags table
func (tags Tags) TableName() string {
	return "Tags"
}

//Challenge structure for Challenge table
type Challenge struct {
	ID           int       `gorm:"column:ID;primary_key;AUTO_INCREMENT" json:"id"`
	Title        string    `gorm:"column:Title" json:"title"`
	Description  string    `gorm:"column:Description" json:"description"`
	Tag          string    `gorm:"column:Tag" json:"tag"`
	VoteCount    int       `gorm:"column:VoteCount" json:"voteCount"`
	CreatedBy    int       `gorm:"column:CreatedBy"`
	CreatedDate  time.Time `gorm:"column:CreatedDate" json:"createdDate"`
	IsDeleted    bool      `gorm:"column:IsDeleted"`
	Collabrators []string  `gorm:"-" json:"collabrators"`
}

//TableName Challenge
func (challenge Challenge) TableName() string {
	return "Challenges"
}

//User structure to store user table
type User struct {
	ID   int    `gorm:"column:ID;primary_key;AUTO_INCREMENT" json:"employeeId"`
	Name string `gorm:"column:Name" json:"employeeName"`
}

//TableName for User
func (user User) TableName() string {
	return "users"
}

//ChallengeCollabration struct
type ChallengeCollabration struct {
	ID          int `gorm:"column:ID;primary_key;AUTO_INCREMENT" json:"id"`
	ChallengeID int `gorm:"column:ChallengeId" json:"challengeId"`
	UserID      int `gorm:"column:UserId" json:"userId"`
}

//TableName for ChallengeCollabration
func (collabration ChallengeCollabration) TableName() string {
	return "ChallengeCollabration"
}
