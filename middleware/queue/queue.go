package queue

import (
	"errors"
	"log"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	usermw "github.com/avidreder/monmach-api/middleware/user"
	"github.com/avidreder/monmach-api/resources/queue"

	"github.com/labstack/echo"
)

// GetUserQueue retieves the user queue from the context
func GetUserQueue(c echo.Context) *queue.Queue {
	return c.Get("queue").(*queue.Queue)
}

// LoadUserQueue places a user into the contest
func LoadUserQueue(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := usermw.GetUser(c)
		userQueue := queue.Queue{}
		store := stmw.GetStore(c)
		err := store.GetByKey("queue", &userQueue, "userid", user.ID)
		if err != nil {
			return errors.New("Could not retrieve user queue")
		}
		log.Printf("user queue: %+v", userQueue)
		c.Set("queue", &userQueue)
		return h(c)
	}
}
