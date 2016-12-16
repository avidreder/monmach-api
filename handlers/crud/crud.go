package crud

import (
	"net/http"

	"github.com/labstack/echo"
)

// Success retrieves an existing user in the store
func Success(c echo.Context) error {
	return c.String(http.StatusOK, "Operation was fine")
}

// Results retrieves an existing user in the store
func Results(c echo.Context) error {
	result := c.Get("result")
	return c.JSON(http.StatusOK, result)
}
