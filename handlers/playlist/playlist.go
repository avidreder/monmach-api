package playlist

import (
	"net/http"

	"github.com/labstack/echo"

	stmw "github.com/avidreder/monmach-api/middleware/store"
)

const tableName = "test"

// Create inserts a new playlist into the store
func Create(c echo.Context) error {
	store := stmw.GetStore(c)
	payload := map[string]interface{}{"name": "Andrew", "age": 30}
	rows, err := store.Create(tableName, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, rows)
}
