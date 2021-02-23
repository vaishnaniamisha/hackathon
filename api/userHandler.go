package api

import (
	"scripbox/hackathon/service"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserService service.UserServiceInyerface
}

func (u *UserHandler) GetUserDetails(c echo.Context) error {
	return nil
}
