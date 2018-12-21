package main

import (
	"github.com/ikasamt/zapp-example/handlers"
	"github.com/ikasamt/zapp-example/models"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ikasamt/zapp/zapp"
	"google.golang.org/appengine"
)

var zappEnvironment zapp.Environment
var isDev = false

const SessionMaxAge = (60 * 60 * 24) * 365 * 5 // 5年

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	zappEnvironments, err := zapp.ReadEnvironments()
	if err != nil {
		log.Println(err)
		return
	}

	zappEnvironment = zappEnvironments[`production`]
	if appengine.IsDevAppServer() {
		zappEnvironment = zappEnvironments[`development`]
		isDev = true
	}

	handlers.ENV = &zappEnvironment
	models.ENV = &zappEnvironment

	sessionSalt := zappEnvironment[`session_salt`].(string)
	sessionStore := cookie.NewStore([]byte(sessionSalt))
	sessionStore.Options(sessions.Options{Path: "/", MaxAge: SessionMaxAge})

	gin.SetMode(gin.ReleaseMode)
	root := gin.Default()

	root.Use(sessions.Sessions(`mysession2`, sessionStore))
	root.Use(zapp.LoggingMiddleware(isDev))

	// ログイン前の画面設定
	root.Use()
	{
		root.GET("/ping", handlers.Pong)
		root.GET("/user/login", handlers.Pong)
		root.GET("/user/logout", handlers.UserLogout)

		root.GET("/dummy", handlers.DummyCreateHandler)

		root.GET("/user/new", models.UserNewHandler)
		root.GET("/user/show/:ID", models.UserShowHandler)
		root.POST("/user/new", models.UserCreateHandler)
	}

	// ログイン後の画面設定
	root.Use(
		models.UserAuthRequired(),
	)
	{
		root.GET("/", handlers.TopIndex)
		// root.GET("/user/new", models.UserNewHandler)
		// root.GET("/user/show/:ID", models.UserShowHandler)
		// root.POST("/user/new", models.UserCreateHandler)
	}

	http.Handle("/", root)
	appengine.Main()
}
