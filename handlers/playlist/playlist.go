package playlist

import (
	"net/http"

	"github.com/labstack/echo"
)

// RetrieveTracks retrieves the processed tracks
func RetrieveTracks(c echo.Context) error {
	tracks := c.Get("tracks")
	return c.JSON(http.StatusOK, tracks)
}
