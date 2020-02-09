package router

import (
	"fmt"
	"mediapedia-api/client"
	"mediapedia-api/handler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

// New connects to a database and starts mediapedia api
func New() *echo.Echo {

	// Get Runtime Params
	mongoServerURL := os.Getenv("MONGO_SERVER_URL")
	mongoUser := os.Getenv("MONGO_USER")
	mongoPass := os.Getenv("MONGO_PASS")

	// create a new echo instance
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)

	//set all middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database connection
	mongo := fmt.Sprintf("mongodb://%s:%s@%s", mongoUser, mongoPass, mongoServerURL)
	db, err := client.NewSession(mongo)
	// db, err := mgo.Dial(mongo)
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

	return e
}
