package api

import (
	"net/http"
	"scripbox/hackathon/service"

	"github.com/labstack/echo"
)

//ChallengeHandler structure
type ChallengeHandler struct {
	ChallengeService service.ChallengeServiceInterface
}

//Ping to check health of server
func (ch *ChallengeHandler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "Status OK")
}

//GetTags to get list of tags
func (ch *ChallengeHandler) GetTags(c echo.Context) error {
	tagList, err := ch.ChallengeService.GetChallengeTagList()
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, tagList)
}
