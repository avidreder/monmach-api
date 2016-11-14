package playlist

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	stmw "github.com/avidreder/monmach-api/middleware/store"
)

const tableName = "test"

type testStruct struct {
	ID   int64
	Name string
	Age  int64
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

// Update updates an existing playlist in the store
func Update(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	store := stmw.GetStore(c)
	payload := map[string]interface{}{"name": "Andrew111", "age": 31}
	rows, err := store.Update(tableName, numId, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, rows)
}

// Get retrieves an existing playlist in the store
func Get(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	testy := testStruct{}
	store := stmw.GetStore(c)
	err = store.Get(tableName, numId, &testy.ID, &testy.Name, &testy.Age)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, testy)
}

// Delete deletes an existing playlist in the store
func Delete(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	store := stmw.GetStore(c)
	err = store.Delete(tableName, numId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Playlist %s deleted", id))
}
