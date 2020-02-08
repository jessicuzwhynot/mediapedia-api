package router

import (
	"mediapedia-api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)

	//set all middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database connection
	db, err := mgo.Dial("localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Create indices
	if err = db.Copy().DB("mediapedia").C("games").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	// Initialize handler
	handler := &handler.Handler{DB: db}

	// Route / to handler function
	e.POST("media/games", handler.AddGame)
	e.GET("media/games/:gameID", handler.GetGame)
	// e.GET("user/:userID/media", handler.GetMedia)
	// e.GET("user/:userID/ratings", handler.GetRatings)
	// e.GET("/main", handler.MainAdmin)

	return e
}
