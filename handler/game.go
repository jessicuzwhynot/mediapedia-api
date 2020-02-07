package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Store struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Game struct {
	CoverArt         string  `json:"coverArt"`
	Description      string  `json:"description"`
	ID               string  `json:"id"`
	MediapediaRating int     `json:"mediapediaRating"`
	Metacritic       int     `json:"metacritic"`
	Publisher        string  `json:"publisher"`
	ReleaseDate      string  `json:"releaseDate"`
	Stores           []Store `json:"stores"`
}

func GetGame(c echo.Context) error {
	// gameID := c.Param("gameID")

	game := &Game{
		CoverArt:         "https://gamecoverart.com",
		Description:      "Lorem Ipsum",
		ID:               "123",
		MediapediaRating: 96,
		Metacritic:       95,
		Publisher:        "Microsoft",
		ReleaseDate:      "2019-12-31",
		Stores: []Store{
			Store{
				Name: "Steam",
				URL:  "https://store.steampowered.com",
			},
			Store{
				Name: "GOG",
				URL:  "https://gog.com",
			},
		},
	}
	return c.JSON(http.StatusOK, game)
}
