package core

import (
	"estj/src/api"
	"estj/src/config"
	"estj/src/router"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// singleton 객체값(pointer)
var app *App

type App struct {
}

func init() {
	initApp()
}

func GetApp() *App {
	if app == nil {
		initApp()
	}
	return app
}

func initApp() {
	app = new(App)
}

func (app *App) RunApp() {
	// Get argument.
	profile := flag.String("profile", "prod", "Check profile option")
	flag.Parse()

	// Init environment variables.
	InitEnvironment(*profile)

	// Set database.
	dbInstance := config.GetDB()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(dbInstance)

	// Set controller list.
	listFunc := []func(engine *gin.Engine){
		api.HealthController,
		api.UserController,
	}

	// Run server.
	route := router.GetRouter()
	route.SetTrustedProxiesPlatforms()
	route.SetCORS()
	route.SetRouting(listFunc)
	route.Start()
}
