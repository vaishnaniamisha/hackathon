package api

import (
	"encoding/json"
	"fmt"
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

//UpvoteChallenge to vote
func (ch *ChallengeHandler) UpvoteChallenge(c echo.Context) error {
	userIDParam := c.Request().Header.Get("UserID")
	userID, err := validateUserID(userIDParam)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	challengeIDParam := c.QueryParam("challengeID")
	challengeID, err := validtaeChallengeID(challengeIDParam)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	challenge, err := ch.ChallengeService.GetChallengeDetails(challengeID)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	err = validateVote(userID, challenge)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	challenge, err = ch.ChallengeService.UpvoteChallenge(challenge)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, challenge)
}

//CollabrateChallenge handler to add challenge collabration
func (ch ChallengeHandler) CollabrateChallenge(c echo.Context) error {
	return nil
}

//GetAllChallenge handler to list all challneges
func (ch ChallengeHandler) GetAllChallenge(c echo.Context) error {
	params := c.QueryParams()
	fmt.Println(params)
	fmt.Println("--", c.QueryParams().Get("sortby"))
	challenges, err := ch.ChallengeService.ListAllChallenges(params)
	if err != nil {
		c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, challenges)
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

func validtaeChallengeID(challengeIDStr string) (int, *errors.ServiceError) {
	if challengeIDStr == "" {
		return 0, &errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: ChallengeID missing",
		}
	}
	challengeID, err := strconv.Atoi(challengeIDStr)
	if err != nil {
		return 0, &errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: ChallengeID is not valid",
		}
	}
	return challengeID, nil
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

func validateVote(userID int, challenge model.Challenge) *errors.ServiceError {
	if challenge.CreatedBy == userID {
		return &errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: Creator can't upvote challenege",
		}
	}

	return nil
}
