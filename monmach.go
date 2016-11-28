package main

import (
	"log"

	authh "github.com/avidreder/monmach-api/handlers/auth"
	gh "github.com/avidreder/monmach-api/handlers/genre"
	plh "github.com/avidreder/monmach-api/handlers/playlist"
	qh "github.com/avidreder/monmach-api/handlers/queue"
	spoth "github.com/avidreder/monmach-api/handlers/spotify"
	th "github.com/avidreder/monmach-api/handlers/track"
	uh "github.com/avidreder/monmach-api/handlers/user"
	authmw "github.com/avidreder/monmach-api/middleware/auth"
	stmw "github.com/avidreder/monmach-api/middleware/store"
	usermw "github.com/avidreder/monmach-api/middleware/user"
	"github.com/avidreder/monmach-api/resources/config"
	spotifyR "github.com/avidreder/monmach-api/resources/spotify"
	"github.com/avidreder/monmach-api/resources/store/postgres"

	"github.com/labstack/echo"
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
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Could not get Service Config: %v", err)
	}
	// Load middleware for all routes
	server.Use(emw.Logger())
	server.Use(emw.Recover())
	server.Use(emw.CORS())
	server.Static("/libs", config.ReactPath+"libs")
	server.Static("/img", config.ReactPath+"img")
	server.File("/bundle.js", config.ReactPath+"bundle.js")

	monmach := server.Group("/")
	monmach.Use(authmw.LoadStore,
		authmw.CheckLogin)
	monmach.Static("", config.ReactPath+"index.html")

	login := server.Group("/login")
	login.Static("", config.ReactPath+"login.html")

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

	// Load routes for spotify
	spotify := server.Group("/spotify")
	spotify.Use(authmw.LoadSpotifyProvider,
		stmw.LoadStore,
		authmw.LoadStore,
		usermw.LoadUser,
		spotifyR.LoadClient)
	spotify.GET("/discover", spoth.DiscoverPlaylist)

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

	// Load routes for users
	users := server.Group("/users")
	users.Use(stmw.LoadStore)
	users.POST("/new", uh.Create)
	users.GET("/:id", uh.Get)
	users.GET("/all", uh.GetAll)
	users.PUT("/:id", uh.Update)
	users.DELETE("/:id", uh.Delete)

	// Load routes for tracks
	tracks := server.Group("/tracks")
	tracks.Use(stmw.LoadStore)
	tracks.POST("/new", th.Create)
	tracks.GET("/:id", th.Get)
	tracks.GET("/all", th.GetAll)
	tracks.PUT("/:id", th.Update)
	tracks.DELETE("/:id", th.Delete)

	// Load routes for queues
	queues := server.Group("/queues")
	queues.Use(stmw.LoadStore)
	queues.POST("/new", qh.Create)
	queues.GET("/:id", qh.Get)
	queues.GET("/all", qh.GetAll)
	queues.PUT("/:id", qh.Update)
	queues.DELETE("/:id", qh.Delete)

	log.Println("Starting...")
	server.Start(":3000")
}
