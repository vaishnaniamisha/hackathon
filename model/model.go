package model

//Tags struct to map Tags db table
type Tags struct {
	Tag string `gorm:"column:Tag"`
}

//TableName for Tags table
func (tags Tags) TableName() string {
	return "Tags"
}
