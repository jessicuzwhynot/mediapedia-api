package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserGame struct {
	ID       string `json:"id"`
	LastSeen string `json:"lastSeen"`
	Owned    bool   `json:"owned"`
	Rating   int    `json:"rating"`
}

type Media struct {
	Games []UserGame
}

func GetMedia(c echo.Context) error {
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
