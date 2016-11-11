package playlist

import (
	"net/http"

	"github.com/labstack/echo"

	stmw "github.com/avidreder/monmach-api/middleware/store"
)

const tableName = "test"

type testStruct struct {
	name string
	age  int64
}

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

// // Update updates an existing playlist in the store
// func Update(c echo.Context) error {
// 	id := c.Params("id")
// 	store := stmw.GetStore(c)
// 	payload := map[string]interface{}{"name": "Andrew", "age": 30}
// 	rows, err := store.Update(tableName, id, payload)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.HTML(http.StatusOK, rows)
// }

// Get retrieves an existing playlist in the store
func Get(c echo.Context) error {
	id := c.Param("id")
	store := stmw.GetStore(c)
	playlist, err := store.Get(tableName, id, testStruct{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, playlist)
}
