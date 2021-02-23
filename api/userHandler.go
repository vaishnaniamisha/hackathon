package api

import (
	"net/http"
	"scripbox/hackathon/errors"
	"scripbox/hackathon/service"
	"strconv"

	"github.com/labstack/echo"
)

//UserHandler structure
type UserHandler struct {
	UserService service.UserServiceInyerface
}

//GetUserDetails handler to validate and get user details
func (u *UserHandler) GetUserDetails(c echo.Context) error {
	userIDParam := c.QueryParam("userID")
	if userIDParam == "" {
		err := errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: Please provide userID",
		}
		return c.JSON(err.Code, err)
	}
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		err := errors.ServiceError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid Input: Please provide valid userID",
		}
		return c.JSON(err.Code, err)
	}
	user, serviceErr := u.UserService.GetUserDetails(userID)
	if serviceErr != nil {
		return c.JSON(serviceErr.Code, serviceErr)
	}
	return c.JSON(http.StatusOK, user)
}
