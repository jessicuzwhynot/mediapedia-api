package main

import (
	"mediapedia-api/router"
)

func main() {
	e := router.New()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
