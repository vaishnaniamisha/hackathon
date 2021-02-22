package main

import (
	"log"
	"scripbox/hackathon/api"
	"scripbox/hackathon/application"
	"scripbox/hackathon/config"
)

func main() {
	cfg, err := application.LoadConfiguration()
	if err != nil {
		log.Fatal("Error reading server configuration : ", err.Error())
	}
	err = RunHTTPServer(cfg.Server)
	if err != nil {
		log.Fatal("Error while running server : ", err.Error())
	}
}

//RunHTTPServer to start server
func RunHTTPServer(cfg *config.ServerConfiguration) error {
	server := api.NewServer(cfg)
	err := server.Start()
	if err != nil {
		return err
	}
	return nil
}
