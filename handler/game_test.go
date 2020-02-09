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
	r := e.Router()
	r.Add("GET", "/:jwt", func(echo.Context) error { return nil })
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/media/games/:gameID")
	c.SetParamNames("gameID")
	c.SetParamValues("5e3f4ddb2c9f1e0001c6403f")

	database := mock.NewMockSession()
	handler := &Handler{DB: database}

	gameJSON := "{\"id\":\"5e3f4ddb2c9f1e0001c6403f\",\"name\":\"Test\",\"coverArt\":\"https://gamecoverart.com\",\"description\":\"Lorem Ipsum\",\"mediapediaRating\":96,\"metacritic\":95,\"publisher\":\"Microsoft\",\"releaseDate\":\"2019-12-31\",\"stores\":[{\"id\":\"\",\"name\":\"Steam\",\"url\":\"https://store.steampowered.com\"},{\"id\":\"\",\"name\":\"GOG\",\"url\":\"https://gog.com\"}]}\n"

	// Assertions
	if assert.NoError(t, handler.GetGame(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, gameJSON, rec.Body.String())
	}
}
