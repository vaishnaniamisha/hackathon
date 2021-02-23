package api_test

import (
	"net/http"
	"net/http/httptest"
	"scripbox/hackathon/api"
	"scripbox/hackathon/errors"
	"scripbox/hackathon/service"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetUserDetails(t *testing.T) {
	tests := []struct {
		name           string
		userID         int
		userIDStr      string
		expectedResult string
		resStatus      int
	}{
		{"success", 1001, "1001", `{"employeeId":"1001","EmployeeName":"User 1"}`, http.StatusOK},
		{"InvalidUserId", 0, "abc", `{"Code":400,"ErrorMessage":"Invalid Input: Please provide valid userID"}`, http.StatusBadRequest},
		{"UserNotExist", 0, "0", `{"Code":400,"ErrorMessage":"UserId does not exist"}`, http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := new(service.MockUserService)
			e := echo.New()
			req := httptest.NewRequest(echo.GET, "/v1/hackathon/user/"+tt.userIDStr, nil)
			rec := httptest.NewRecorder()
			context := e.NewContext(req, rec)
			handler := api.UserHandler{
				UserService: mockUserService,
			}
			if tt.name == "UserNotExist" {
				mockUserService.On("GetUserDetails", tt.userID).Return(&errors.ServiceError{
					Code:         http.StatusNotFound,
					ErrorMessage: "UserId does not exist",
				})
			} else {
				mockUserService.On("GetUserDetails", tt.userID).Return((*errors.ServiceError)(nil))
			}

			err := handler.GetUserDetails(context)
			assert.Nil(t, err)
			assert.Equal(t, tt.resStatus, rec.Code)

			resStr := rec.Body.String()
			assert.Equal(t, tt.expectedResult, resStr[:len(resStr)-1])
		})
	}
}
