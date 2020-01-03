package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, " You are on the Admin Page !!!")

}
