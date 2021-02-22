package application

import (
	"fmt"
	"scripbox/hackathon/config"
	"scripbox/hackathon/lib/database"

	"github.com/spf13/viper"
)

//LoadConfiguration to load configuration file
func LoadConfiguration() (*config.Configuration, error) {
	var cfg *config.Configuration
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return nil, err
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}
	database.DbConfig = *cfg.Database
	return cfg, nil
}
