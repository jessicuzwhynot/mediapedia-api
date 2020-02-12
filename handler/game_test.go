package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetGame(t *testing.T) {
	// Setup
	e := echo.New()

	// dirty hack cause echo is broke! https://github.com/labstack/echo/pull/1498#issuecomment-582799854
	r := e.Router()
	r.Add("GET", "/:jwt", func(echo.Context) error { return nil })

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/media/games/:gameID")
	c.SetParamNames("gameID")
	c.SetParamValues("1")

	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows([]string{"id", "title", "cover_art", "description", "mediapedia_rating", "metacritic_rating", "publisher", "release_date"}).
		AddRow("1", "Bloodborne", "https://gamecoverart.com", "LoremIpsum", "96", "95", "FromSoftware", "2019-12-31T00:00:00Z")

	mock.ExpectQuery("SELECT id, title, cover_art, description, mediapedia_rating, metacritic_rating, publisher, release_date FROM game WHERE id=?").
		WithArgs("1").
		WillReturnRows(rows)

	handler := &Handler{DB: database}

	gameJSON := "{\"id\":\"1\",\"title\":\"Bloodborne\",\"coverArt\":\"https://gamecoverart.com\",\"description\":\"LoremIpsum\",\"mediapediaRating\":96,\"metacritic\":95,\"publisher\":\"FromSoftware\",\"releaseDate\":\"2019-12-31T00:00:00Z\"}\n"

	// Assertions
	if assert.NoError(t, handler.GetGame(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, gameJSON, rec.Body.String())
	}
}
