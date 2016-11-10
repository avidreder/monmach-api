package main

import (
	"log"

	plh "github.com/avidreder/monmach-api/handlers/playlist"
	// yh "github.com/avidreder/monmach-api/handlers/youtube"
	// plmw "github.com/avidreder/monmach-api/middleware/playlist"
	"github.com/avidreder/monmach-api/resources/store/postgres"
	// "github.com/avidreder/monmach-api/resources/shows/mongo"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	emw "github.com/labstack/echo/middleware"
)

func main() {
	server := echo.New()
	store, _ := postgres.Get()
	err := store.Connect()
	if err != nil {
		log.Printf("Could not connect to Postgres: %v", err)
	} else {
		log.Print("Connected to Postgres")
	}
	// Load middleware for all routes
	server.Use(emw.Logger())
	server.Use(emw.Recover())
	server.Use(emw.CORS())

	// Load routes for playlists
	bitshows := server.Group("/playlists")
	bitshows.GET("", plh.GetShows)

	log.Println("Starting...")
	server.Run(standard.New(":3000"))
}
