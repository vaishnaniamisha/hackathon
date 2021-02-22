package database

import (
	"fmt"
	"log"
	"scripbox/hackathon/config"

	"github.com/jinzhu/gorm"
	//pq is postgresql driver
	_ "github.com/lib/pq"
)

//DBClient to implement db methods
type DBClient struct {
	GormDB *gorm.DB
}

//DbConfig to store database configuration
var DbConfig config.DBConfiguration

//DBConnect to initiate DB connection
func (dc *DBClient) DBConnect() error {
	var err error
	// Format DB configs to required format connect DB
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DbConfig.DBHost, DbConfig.DBUserName, DbConfig.DBPassword, DbConfig.DBName, DbConfig.DBPort)
	dc.GormDB, err = gorm.Open("postgres", dbinfo)

	if err != nil {
		log.Println("Unable to connect DB")
		return err
	}
	err = dc.GormDB.DB().Ping()
	if err != nil {
		log.Println("Db Ping failed: ", err.Error())
	}
	log.Printf("Postgres started at %s PORT", DbConfig.DBPort)
	return err
}
