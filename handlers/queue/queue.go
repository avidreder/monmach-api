package queue

import (
	"net/http"

	"github.com/labstack/echo"
)

// RetrieveQueue retrieves the user's queue
func RetrieveQueue(c echo.Context) error {
	queue := c.Get("queue")
	return c.JSON(http.StatusOK, queue)
}
