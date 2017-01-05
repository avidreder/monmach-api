package user

import (
	"fmt"
	"log"
	"net/http"

	authmw "github.com/avidreder/monmach-api/middleware/auth"
	stmw "github.com/avidreder/monmach-api/middleware/store"
	"github.com/avidreder/monmach-api/resources/user"

	"github.com/labstack/echo"
)

// GetSpotifyProvider retieves provider from the context
func GetUser(c echo.Context) *user.User {
	return c.Get("user").(*user.User)
}

// LoadUser places a user into the contest
func LoadUser(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionStore := authmw.GetStore(c)
		session, err := sessionStore.Get(c.Request(), "auth-session")
		if session.IsNew || err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Could not retrieve logged-in user: %v", err))
		}
		userEmail := session.Values["email"].(string)
		log.Printf("Loading User: %s", userEmail)
		user := user.User{}
		store := stmw.GetStore(c)
		err = store.GetByKey("users", &user, "Email", userEmail)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Could not retrieve logged-in user: %v", err))
		}
		c.Set("user", &user)
		return h(c)
	}
}
