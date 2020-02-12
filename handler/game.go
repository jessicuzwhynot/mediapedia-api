package handler

import (
	"net/http"

	"mediapedia-api/model"

	"github.com/labstack/echo/v4"
)

// func (h *Handler) AddGame(c echo.Context) (err error) {

// 	game := &model.Game{
// 		Title:            "Test",
// 		CoverArt:         "https://gamecoverart.com",
// 		Description:      "Lorem Ipsum",
// 		MediapediaRating: 96,
// 		MetacriticRating: 95,
// 		Publisher:        "Microsoft",
// 		ReleaseDate:      "2019-12-31",
// 		// Stores: []model.Store{
// 		// 	model.Store{
// 		// 		Name: "Steam",
// 		// 		URL:  "https://store.steampowered.com",
// 		// 	},
// 		// 	model.Store{
// 		// 		Name: "GOG",
// 		// 		URL:  "https://gog.com",
// 		// 	},
// 		// },
// 	}

// 	if err = c.Bind(game); err != nil {
// 		return
// 	}

// 	// db := h.DB.Clone()
// 	// defer db.Close()
// 	// if err = db.DB("mediapedia").C("games").Insert(game); err != nil {
// 	// 	return
// 	// }

// 	return c.JSON(http.StatusOK, game)
// }

// GetGame retrieves corresponding game from db
func (h *Handler) GetGame(c echo.Context) (err error) {
	gameID := c.Param("gameID")

	game := &model.Game{}

	sql := "SELECT id, title, cover_art, description, mediapedia_rating, metacritic_rating, publisher, release_date FROM game WHERE id=$1"
	result := h.DB.QueryRow(sql, gameID)
	err = result.Scan(&game.ID, &game.Title, &game.CoverArt, &game.Description, &game.MediapediaRating, &game.MetacriticRating, &game.Publisher, &game.ReleaseDate)

	if err != nil {
		return c.JSON(http.StatusNotFound, game)
	}

	return c.JSON(http.StatusOK, game)
}
