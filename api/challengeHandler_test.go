package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"scripbox/hackathon/api"
	"scripbox/hackathon/errors"
	"scripbox/hackathon/model"
	"scripbox/hackathon/service"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/v1/hackathon/ping", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	handler := api.ChallengeHandler{
		ChallengeService: service.MockChallengeService{},
	}
	err := handler.Ping(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetTagsSuccess(t *testing.T) {
	mockChallengeService := new(service.MockChallengeService)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/v1/hackathon/tags", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	handler := api.ChallengeHandler{
		ChallengeService: mockChallengeService,
	}
	tagList := []string{"tag1", "tag2"}
	mockChallengeService.On("GetChallengeTagList").Return(tagList, (*errors.ServiceError)(nil))
	err := handler.GetTags(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body)
}

func TestGetTagsError(t *testing.T) {
	mockChallengeService := new(service.MockChallengeService)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/v1/hackathon/tags", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	handler := api.ChallengeHandler{
		ChallengeService: mockChallengeService,
	}
	tagList := []string{"tag1", "tag2"}
	mockChallengeService.On("GetChallengeTagList").Return(tagList, &errors.ServiceError{Code: http.StatusInternalServerError, ErrorMessage: "Internal Server Error"})
	err := handler.GetTags(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestCreateChallenge(t *testing.T) {
	tests := []struct {
		name      string
		body      model.Challenge
		userID    string
		want      string
		resStatus int
	}{
		{"successTest", model.Challenge{
			Title:       "Challenge 1",
			Description: "Challenge 1 Description",
			Tag:         "Tag 1",
		}, "1001", `"Success"`, http.StatusOK},
		{"InvalidBody", model.Challenge{},
			"1001", `{"Code":400,"ErrorMessage":"Invalid Input: Please provide valid input"}`, http.StatusBadRequest},
		{"EmptyTitle", model.Challenge{
			Title:       "",
			Description: "Challenge 2 Description",
			Tag:         "Tag 2",
		}, "1001", `{"Code":400,"ErrorMessage":"Invalid Input: Please provide title"}`, http.StatusBadRequest},
		{"EmptyDescription", model.Challenge{
			Title:       "Challenge 2",
			Description: "",
			Tag:         "Tag 2",
		}, "1001", `{"Code":400,"ErrorMessage":"Invalid Input: Please provide description"}`, http.StatusBadRequest},
		{"InvalidTag", model.Challenge{
			Title:       "Challenge 3",
			Description: "Challenge 3 Description",
			Tag:         "Dummy tag",
		}, "1001", `{"Code":400,"ErrorMessage":"Invalid Input: Please provide valid tag"}`, http.StatusBadRequest},
		{"EmptyUserId", model.Challenge{
			Title:       "Challenge 3",
			Description: "Challenge 3 Description",
			Tag:         "test tag",
		}, "", `{"Code":400,"ErrorMessage":"Invalid Input: UserId missing"}`, http.StatusBadRequest},
		{"InvalidUserId", model.Challenge{
			Title:       "Challenge 3",
			Description: "Challenge 3 Description",
			Tag:         "test tag",
		}, "abc", `{"Code":400,"ErrorMessage":"Invalid Input: UserId is not valid"}`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockChallengeService := new(service.MockChallengeService)
			jsonbody := []byte{}
			e := echo.New()
			if tt.name != "InvalidBody" {
				jsonbody, _ = json.Marshal(tt.body)
			}
			req := httptest.NewRequest(echo.POST, "/v1/hackathon/challenge", strings.NewReader(string(jsonbody)))
			rec := httptest.NewRecorder()
			req.Header.Add("UserID", tt.userID)
			context := e.NewContext(req, rec)
			handler := api.ChallengeHandler{
				ChallengeService: mockChallengeService,
			}
			if tt.name == "InvalidTag" {
				mockChallengeService.On("ValidateTag", tt.body.Tag).Return(&errors.ServiceError{
					Code:         http.StatusBadRequest,
					ErrorMessage: "Invalid Input: Please provide valid tag",
				})
			} else {
				mockChallengeService.On("ValidateTag", tt.body.Tag).Return((*errors.ServiceError)(nil))
			}
			mockChallengeService.On("AddChallenge", tt.body).Return((*errors.ServiceError)(nil))
			err := handler.CreateChallenge(context)
			assert.Nil(t, err)
			assert.Equal(t, tt.resStatus, rec.Code)
			resStr := rec.Body.String()
			assert.Equal(t, tt.want, resStr[:len(resStr)-1])
		})
	}

}

func TestCreateChallengeError(t *testing.T) {
	mockChallengeService := new(service.MockChallengeService)
	challenge := model.Challenge{
		Title:       "Challenge 1",
		Description: "Challenge 1 Description",
		Tag:         "Tag 1",
	}
	e := echo.New()
	jsonbody, _ := json.Marshal(challenge)
	req := httptest.NewRequest(echo.POST, "/v1/hackathon/challenge", strings.NewReader(string(jsonbody)))
	rec := httptest.NewRecorder()
	req.Header.Add("UserID", "1001")
	context := e.NewContext(req, rec)
	handler := api.ChallengeHandler{
		ChallengeService: mockChallengeService,
	}
	mockChallengeService.On("AddChallenge", challenge).Return(&errors.ServiceError{Code: http.StatusInternalServerError, ErrorMessage: "Internal server error"})
	err := handler.CreateChallenge(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
