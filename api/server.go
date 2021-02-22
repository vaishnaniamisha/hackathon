package api

import (
	"fmt"
	"log"
	"scripbox/hackathon/config"
	"scripbox/hackathon/lib/database"
	"scripbox/hackathon/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Server structure
type Server struct {
	ServerConfiguration *config.ServerConfiguration
	router              *echo.Echo
}

//NewServer to configure new server
func NewServer(cfg *config.ServerConfiguration) *Server {
	server := &Server{
		ServerConfiguration: cfg,
		router:              echo.New(),
	}
	server.router.Use(middleware.Recover())
	server.router.Use(middleware.CORS())

	server.registerAPIs()
	return server
}

func (s *Server) registerAPIs() {

	_dbrepo, err := initDBClient()
	if err != nil {
		log.Fatal("Error Initializing Db instance")
	}
	handler := ChallengeHandler{
		ChallengeService: service.ChallengeService{
			DbClient: _dbrepo,
		},
	}

	group := s.router.Group("/v1")
	group.GET("/hackathon/ping", handler.Ping)
	group.GET("/hackathon/tags", handler.GetTags)

}

//Start to start server
func (s *Server) Start() error {
	err := s.router.Start(":" + s.ServerConfiguration.Port)
	return err
}

//initDBClient to initialize database connection
func initDBClient() (database.DBClientInterface, error) {
	_dbrepo := &database.DBClient{}
	err := _dbrepo.DBConnect()
	if err != nil {
		return _dbrepo, err
	}
	fmt.Println("connection to the database was successful")
	return _dbrepo, nil
}
