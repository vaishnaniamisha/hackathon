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
