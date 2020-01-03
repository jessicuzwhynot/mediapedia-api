package router

import (
	"github.com/HammerMeetNail/mediapedia-api/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	admin := e.Group("/admin")

	//set all middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//set main routes
	mainGroup(e)

	//set groupRoutes
	adminGroup(admin)
	return e
}

func mainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("media/games/:gameID", handler.GetGame)
	e.GET("user/:userID/media", handler.GetMedia)
	e.GET("user/:userID/ratings", handler.GetRatings)
}

func adminGroup(g *echo.Group) {
	g.GET("/main", handler.MainAdmin)
}
