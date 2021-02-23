package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"scripbox/hackathon/api"
	"scripbox/hackathon/errors"
	"scripbox/hackathon/model"
	"scripbox/hackathon/service"
	"strconv"
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
			CreatedBy:   1001,
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
		CreatedBy:   1001,
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
	mockChallengeService.On("ValidateTag", challenge.Tag).Return((*errors.ServiceError)(nil))
	mockChallengeService.On("AddChallenge", challenge).Return(&errors.ServiceError{Code: http.StatusInternalServerError, ErrorMessage: "Internal server error"})
	err := handler.CreateChallenge(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestUpvoteChallenge(t *testing.T) {
	challenge := model.Challenge{
		ID:          1001,
		CreatedBy:   1001,
		VoteCount:   0,
		Title:       "Challenge 1",
		Description: "Challenge 1 Description",
	}
	result := challenge
	result.VoteCount = result.VoteCount + 1
	challengeJSON, _ := json.Marshal(result)
	tests := []struct {
		name            string
		userID          int
		challengeID     int
		challenge       model.Challenge
		resultChallenge model.Challenge
		want            string
		resStatus       int
	}{
		{
			"Success", 1002, 1001, challenge, result, string(challengeJSON), http.StatusOK,
		},
		{
			"InvalidUser", 1001, 1001, challenge, result, `{"Code":400,"ErrorMessage":"Invalid Input: Creator can't upvote challenege"}`, http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockChallengeService := new(service.MockChallengeService)
			e := echo.New()

			req := httptest.NewRequest(echo.PUT, "/v1/hackathon/vote?challengeID="+strconv.Itoa(tt.challengeID), nil)
			rec := httptest.NewRecorder()
			req.Header.Add("UserID", strconv.Itoa(tt.userID))

			context := e.NewContext(req, rec)
			handler := api.ChallengeHandler{
				ChallengeService: mockChallengeService,
			}
			mockChallengeService.On("GetChallengeDetails", tt.challengeID).Return(tt.challenge, (*errors.ServiceError)(nil))
			mockChallengeService.On("UpvoteChallenge", tt.challenge).Return(tt.resultChallenge, (*errors.ServiceError)(nil))

			err := handler.UpvoteChallenge(context)
			assert.Nil(t, err)
			assert.Equal(t, tt.resStatus, rec.Code)
			resStr := rec.Body.String()
			assert.Equal(t, tt.want, resStr[:len(resStr)-1])
		})
	}
}

func TestGetAllChallenge(t *testing.T) {
	challenges := []model.Challenge{
		{
			ID:          1001,
			CreatedBy:   1001,
			VoteCount:   2,
			Title:       "Challenge 1",
			Description: "Challenge 1 Description",
		},
		{
			ID:          1002,
			CreatedBy:   1001,
			VoteCount:   0,
			Title:       "Challenge 2",
			Description: "Challenge 2 Description",
		},
	}
	challengeJSON, _ := json.Marshal(challenges)
	tests := []struct {
		name           string
		expectedResult string
		challenges     []model.Challenge
		params         string
		resStatus      int
	}{
		{
			"Success", string(challengeJSON), challenges, "voteCount", http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockChallengeService := new(service.MockChallengeService)
			e := echo.New()

			req := httptest.NewRequest(echo.GET, "/v1/hackathon/challenge?sortby="+tt.params, nil)
			rec := httptest.NewRecorder()
			context := e.NewContext(req, rec)
			handler := api.ChallengeHandler{
				ChallengeService: mockChallengeService,
			}
			params := make(map[string][]string)
			params["sortby"] = []string{tt.params}

			mockChallengeService.On("ListAllChallenges", params).Return(tt.challenges, (*errors.ServiceError)(nil))

			err := handler.GetAllChallenge(context)
			assert.Nil(t, err)
			assert.Equal(t, tt.resStatus, rec.Code)
			resStr := rec.Body.String()
			assert.Equal(t, tt.expectedResult, resStr[:len(resStr)-1])
		})
	}
}

func TestGetAllChallengeError(t *testing.T) {
	mockChallengeService := new(service.MockChallengeService)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/v1/hackathon/challenge", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	handler := api.ChallengeHandler{
		ChallengeService: mockChallengeService,
	}
	params := make(map[string][]string)

	mockChallengeService.On("ListAllChallenges", params).Return([]model.Challenge{}, &errors.ServiceError{
		Code:         http.StatusInternalServerError,
		ErrorMessage: "Internal server error",
	})

	err := handler.GetAllChallenge(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

}

func TestCollabrateChallenge(t *testing.T) {
	tests := []struct {
		name           string
		expectedResult string
		userID         int
		challengeID    int
		resStatus      int
	}{
		{
			"Success", `"Success"`, 1001, 1002, http.StatusOK,
		},
		{
			"ServerError", `{"Code":500,"ErrorMessage":"Internal server error"}`, 1001, 1002, http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockChallengeService := new(service.MockChallengeService)
			e := echo.New()

			req := httptest.NewRequest(echo.PUT, "/v1/hackathon/collabration?userID="+strconv.Itoa(tt.userID)+"&challengeID="+strconv.Itoa(tt.userID), nil)
			rec := httptest.NewRecorder()
			context := e.NewContext(req, rec)
			handler := api.ChallengeHandler{
				ChallengeService: mockChallengeService,
			}
			if tt.name == "ServerError" {
				mockChallengeService.On("UpdateCollabration", tt.userID, tt.challengeID).Return(&errors.ServiceError{
					Code:         http.StatusInternalServerError,
					ErrorMessage: "Internal server error",
				})

			} else {
				mockChallengeService.On("UpdateCollabration", tt.userID, tt.challengeID).Return((*errors.ServiceError)(nil))

			}

			err := handler.CollabrateChallenge(context)
			assert.Nil(t, err)
			assert.Equal(t, tt.resStatus, rec.Code)
			resStr := rec.Body.String()
			assert.Equal(t, tt.expectedResult, resStr[:len(resStr)-1])
		})
	}
}
