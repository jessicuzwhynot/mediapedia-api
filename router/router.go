package router

import (
	"database/sql"
	"fmt"
	"mediapedia-api/handler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	// "gopkg.in/mgo.v2"
)

func New() *echo.Echo {

	// Get Runtime Params
	postgresServer := os.Getenv("POSTGRES_SERVER")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPass := os.Getenv("POSTGRES_PASS")
	postgresDB := os.Getenv("POSTGRES_DB")

	// create a new echo instance
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)

	//set all middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Postgress Database connection
	postgres := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", postgresServer, postgresPort, postgresUser, postgresPass, postgresDB)
	db, err := sql.Open(`postgres`, postgres)
	if err != nil {
		panic(err)
	}

	// Check Postgres DB Connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	// Initialize handler
	handler := &handler.Handler{DB: db}

	// Route / to handler function
	// e.POST("media/games", handler.AddGame)
	e.GET("media/games/:gameID", handler.GetGame)
	// e.GET("user/:userID/media", handler.GetMedia)
	// e.GET("user/:userID/ratings", handler.GetRatings)

	return e
}
