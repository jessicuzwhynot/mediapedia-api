package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

type UserGame struct {
	ID       string `json:"id"`
	LastSeen string `json:"lastSeen"`
	Owned    bool   `json:"owned"`
	Rating   int    `json:"rating"`
}

type Media struct {
	Games []UserGame
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("media/games/:gameID", getGame)
	e.GET("user/:userID/media", getMedia)
	e.GET("user/:userID/ratings", getRatings)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func getGame(c echo.Context) error {
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

func getMedia(c echo.Context) error {
	// userID := c.Param("userID")

	media := &Media{
		Games: []UserGame{
			UserGame{
				ID:       "123",
				LastSeen: "2019-12-31",
				Owned:    true,
				Rating:   90,
			},
		},
	}

	return c.JSON(http.StatusOK, media)
}

func getRatings(c echo.Context) error {
	userID := c.Param("userID")
	return c.JSON(http.StatusOK, "userID: "+userID)
}
