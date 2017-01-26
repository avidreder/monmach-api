package main

import (
	"fmt"
	"log"

	authh "github.com/avidreder/monmach-api/handlers/auth"
	crudh "github.com/avidreder/monmach-api/handlers/crud"
	queueh "github.com/avidreder/monmach-api/handlers/queue"
	spoth "github.com/avidreder/monmach-api/handlers/spotify"
	authmw "github.com/avidreder/monmach-api/middleware/auth"
	crudmw "github.com/avidreder/monmach-api/middleware/crud"
	queuemw "github.com/avidreder/monmach-api/middleware/queue"
	spotifymw "github.com/avidreder/monmach-api/middleware/spotify"
	stmw "github.com/avidreder/monmach-api/middleware/store"
	usermw "github.com/avidreder/monmach-api/middleware/user"
	configR "github.com/avidreder/monmach-api/resources/config"
	spotifyR "github.com/avidreder/monmach-api/resources/spotify"
	"github.com/avidreder/monmach-api/resources/store/mongo"

	"github.com/labstack/echo"
	emw "github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
)

const dbURL = "mongodb://localhost:27017"
const db = "monmach"

func main() {
	server := echo.New()
	config, err := configR.GetConfig()
	if err != nil {
		log.Printf("Config error: %v", err)
		panic(err)
	}
	err = mongo.LoadCredentials(config.MongoCredentialsPath)
	if err != nil {
		log.Printf("Config error: %v", err)
		panic(err)
	}
	store, _ := mongo.Get()
	err = store.Connect()
	if err != nil {
		log.Printf("Could not connect to Mongo: %v", err)
	} else {
		log.Print("Connected to Mongo")
	}
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%s@localhost:27017/test", mongo.CurrentCredentials.Username, mongo.CurrentCredentials.Password))
	if err != nil {
		panic(err)
	}
	testData := struct {
		name string
	}{
		name: "Andrew",
	}
	err = session.DB("test").C("test").Insert(&testData)
	if err != nil {
		panic(err)
	}

	err = spotifyR.InitializeSpotifyProvider()
	if err != nil {
		log.Printf("Could not Initialize Spotify: %s", err)
		panic(err)
	}

	// Load middleware for all routes
	server.Use(emw.Logger())
	server.Use(emw.Recover())
	server.Use(emw.CORSWithConfig(emw.CORSConfig{
		AllowOrigins:     []string{config.ClientAddress},
		AllowCredentials: true,
	}))

	logout := server.Group("/logout")
	logout.Use(authmw.LoadStore)
	logout.GET("", authh.LogoutUser)

	getUser := server.Group("/getuser")
	getUser.Use(authmw.LoadStore)
	getUser.GET("", authh.GetUser)

	test := server.Group("/test")
	test.Use(authmw.LoadStore,
		authmw.CheckLogin)
	test.GET("", func(c echo.Context) error {
		return c.HTML(200, "Logged In")
	})
	// Load routes for auth
	auth := server.Group("/auth")
	auth.Use(authmw.LoadSpotifyProvider,
		stmw.LoadStore,
		authmw.LoadStore)
	auth.GET("/spotify", authh.StartAuth)
	auth.GET("/spotify/callback", authh.FinishAuth)

	// Load routes for user queue
	queue := server.Group("/queue")
	queue.Use(authmw.LoadSpotifyProvider,
		stmw.LoadStore,
		authmw.LoadStore,
		usermw.LoadUser,
		spotifymw.LoadClient)
	queue.GET("/user", queueh.RetrieveQueue, queuemw.LoadUserQueue)
	queue.GET("/:playlist", queueh.RetrieveQueue, queuemw.QueueFromPlaylist)

	// Load routes for spotify
	spotify := server.Group("/spotify")
	spotify.Use(authmw.LoadSpotifyProvider,
		stmw.LoadStore,
		authmw.LoadStore,
		usermw.LoadUser,
		queuemw.LoadUserQueue,
		spotifymw.LoadClient)
	spotify.GET("/discover", spoth.DiscoverPlaylist)
	spotify.GET("/playlists", spoth.UserPlaylists)

	// Load routes for crud
	crud := server.Group("/crud/:table")
	crud.Use(stmw.LoadStore)
	crud.POST("/new", crudh.Success, crudmw.Create)
	crud.GET("/:id", crudh.Results, crudmw.Get)
	crud.GET("/all", crudh.Results, crudmw.GetAll)
	crud.PUT("/:id", crudh.Success, crudmw.Update)
	crud.DELETE("/:id", crudh.Success, crudmw.Delete)

	log.Println("Starting...")
	server.Start(":3000")
}
