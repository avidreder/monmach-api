package shows

import (
	"net/http"
	
	"github.com/labstack/echo"
)

// GetBITShows requests from BandsInTown and returns them
func GetShows(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Get("ShowList"))
}
