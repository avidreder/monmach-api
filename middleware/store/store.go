package store

import (
	"fmt"
	"net/http"

	"github.com/avidreder/monmach-api/resources/store/postgres"

	"github.com/labstack/echo"
)

// LoadStore places a data store in the context for later use
func LoadStore(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		dataStore, err := postgres.Get()
		if err != nil {
			errorMessage := fmt.Sprintf("Could not load store into context: %s", err)
			return echo.NewHTTPError(http.StatusUnauthorized, errorMessage)
		}
		c.Set("store", dataStore)
		return h(c)
	}

}

// GetStore retieves a data store from the context
func GetStore(c echo.Context) *postgres.Store {
	return c.Get("store").(*postgres.Store)
}
