package main

import (
	"log"

	sh "github.com/avidreder/show-hawk-server/handlers/shows"
	yh "github.com/avidreder/show-hawk-server/handlers/youtube"
	bitmw "github.com/avidreder/show-hawk-server/middleware/bitshows"
	ytmw "github.com/avidreder/show-hawk-server/middleware/youtube"
	// "github.com/avidreder/show-hawk-server/resources/shows/mongo"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	emw "github.com/labstack/echo/middleware"
)

func main() {
	server := echo.New()
	// _, err := mongo.Connect()
	// if err != nil {
	// 	log.Printf("Could not connect to MongoDB: %v", err)
	// } else {
	// 	log.Print("Connected to MongoDB")
	// }
	// Load middleware for all routes
	server.Use(emw.Logger())
	server.Use(emw.Recover())
	server.Use(emw.CORS())

	// Load routes for Youtube search
	youtube := server.Group("/youtube")
	youtube.Use(ytmw.CompileVideos)
	youtube.GET("", yh.GetVideos)

	// Load routes for BIT shows
	bitshows := server.Group("/bitshows")
	bitshows.Use(bitmw.CompileShows)
	bitshows.GET("", sh.GetShows)

	log.Println("Starting...")
	server.Run(standard.New(":3000"))
}
