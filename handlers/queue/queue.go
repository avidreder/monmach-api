package queue

import (
	"net/http"

	"github.com/labstack/echo"
)

// UserQueue retrieves the user's queue
func UserQueue(c echo.Context) error {
	queue := c.Get("queue")
	return c.JSON(http.StatusOK, queue)
}
