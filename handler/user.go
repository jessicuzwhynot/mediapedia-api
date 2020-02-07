package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRatings(c echo.Context) error {
	userID := c.Param("userID")
	return c.JSON(http.StatusOK, "userID: "+userID)
}
