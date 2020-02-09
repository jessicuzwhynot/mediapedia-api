package handler

import (
	"mediapedia-api/database"
)

type (
	Handler struct {
		DB database.Session
	}
)
