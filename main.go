package main

import (
	"github.com/HammerMeetNail/mediapedia-api/router"
)

func main() {
	e := router.New()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
