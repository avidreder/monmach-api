package main

import (
	"log"

	gh "github.com/avidreder/monmach-api/handlers/genre"
	plh "github.com/avidreder/monmach-api/handlers/playlist"
	// yh "github.com/avidreder/monmach-api/handlers/youtube"
	stmw "github.com/avidreder/monmach-api/middleware/store"
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
	playlists := server.Group("/playlists")
	playlists.Use(stmw.LoadStore)
	playlists.POST("/new", plh.Create)
	playlists.GET("/:id", plh.Get)
	playlists.GET("/all", plh.GetAll)
	playlists.PUT("/:id", plh.Update)
	playlists.DELETE("/:id", plh.Delete)

	// Load routes for genres
	genres := server.Group("/genres")
	genres.Use(stmw.LoadStore)
	genres.POST("/new", gh.Create)
	genres.GET("/:id", gh.Get)
	genres.GET("/all", gh.GetAll)
	genres.PUT("/:id", gh.Update)
	genres.DELETE("/:id", gh.Delete)

	log.Println("Starting...")
	server.Run(standard.New(":3000"))
}
