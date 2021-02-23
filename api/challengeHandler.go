package api

import (
	"encoding/json"
	"net/http"
	"scripbox/hackathon/errors"
	"scripbox/hackathon/model"
	"scripbox/hackathon/service"
	"strconv"

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

//CreateChallenge to add new challege
func (ch *ChallengeHandler) CreateChallenge(c echo.Context) error {
	challenge := model.Challenge{}

	userIDStr := c.Request().Header.Get("UserID")
	userID, serviceErr := validateUserID(userIDStr)
	if serviceErr != nil {
		return c.JSON(serviceErr.Code, serviceErr)
	}

	err := json.NewDecoder(c.Request().Body).Decode(&challenge)
	if err != nil {
		customErr := errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: Please provide valid input",
		}
		return c.JSON(http.StatusBadRequest, customErr)
	}

	serviceErr = ch.validateChallenge(challenge)
	if serviceErr != nil {
		return c.JSON(serviceErr.Code, serviceErr)
	}
	challenge.CreatedBy = userID
	serviceErr = ch.ChallengeService.AddChallenge(challenge)
	if serviceErr != nil {
		return c.JSON(serviceErr.Code, serviceErr)
	}
	return c.JSON(http.StatusOK, "Success")
}

func validateUserID(userID string) (int, *errors.ServiceError) {
	if userID == "" {
		return 0, &errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: UserId missing",
		}
	}
	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		return 0, &errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: UserId is not valid",
		}
	}
	return userIDint, nil
}

func (ch *ChallengeHandler) validateChallenge(challenge model.Challenge) *errors.ServiceError {
	if challenge.Title == "" {
		return &errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: Please provide title",
		}
	}
	if challenge.Description == "" {
		return &errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: Please provide description",
		}
	}
	if challenge.Tag != "" {
		return ch.ChallengeService.ValidateTag(challenge.Tag)
	}
	return nil
}
