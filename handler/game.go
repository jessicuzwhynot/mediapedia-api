package handler

import (
	"net/http"

	"mediapedia-api/model"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) AddGame(c echo.Context) (err error) {

	game := &model.Game{
		ID:               bson.NewObjectId(),
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
		},
	}

	if err = c.Bind(game); err != nil {
		return
	}

	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("mediapedia").C("games").Insert(game); err != nil {
		return
	}

	return c.JSON(http.StatusOK, game)
}

func (h *Handler) GetGame(c echo.Context) (err error) {
	gameID := c.Param("gameID")

	game := model.Game{}
	db := h.DB.Clone()

	if err = db.DB("mediapedia").C("games").FindId(gameID).One(&game); err != nil {
		return

	}

	defer db.Close()

	return c.JSON(http.StatusOK, game)
}
