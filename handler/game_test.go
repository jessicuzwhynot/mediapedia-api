package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"mediapedia-api/mock"
	"mediapedia-api/model"
)

var (
	mockDB = map[string]*model.Game{
		"123": &model.Game{ID: "123",
			Name:             "Test",
			CoverArt:         "https://gamecoverart.com",
			Description:      "Lorem Ipsum",
			MediapediaRating: 96,
			Metacritic:       95,
			Publisher:        "Microsoft",
			ReleaseDate:      "2019-12-31",
			Stores: []model.Store{
				model.Store{
					Name: "Steam",
					URL:  "https://store.steampowered.com",
				},
				model.Store{
					Name: "GOG",
					URL:  "https://gog.com",
				},
			}},
	}
)

func TestGetGame(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/media/games/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	database := mock.NewMockSession()
	handler := &Handler{DB: database}

	// Assertions
	if assert.NoError(t, handler.GetGame(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.Equal(t, userJSON, rec.Body.String())
	}
}
