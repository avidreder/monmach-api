package youtube

import (
	"net/http"
	
	"github.com/labstack/echo"
)

// GetVideos searches Youtube and returns videos
func GetVideos(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Get("VideoList"))
}
