package core

import (
	"estj/src/api"
	"estj/src/config"
	"estj/src/core/sysconfig"
	"estj/src/exception"
	"estj/src/logger"
	log "estj/src/logger"
	"estj/src/router"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"reflect"
)

// singleton 객체값(pointer)
var app *App

type App struct {
}

func init() {
	initApp()
}

func GetApp() *App {
	initApp()
	return app
}

func initApp() {
	if app == nil {
		app = new(App)
	}
}

func (app *App) RunApp(isRunApp bool) {
	if isRunApp {
		// Get argument.
		profile := flag.String("profile", "prod", "Check profile option")
		flag.Parse()

		// Init environment variables.
		sysconfig.InitEnvVariables(*profile)
	} else {
		// Init environment variables.
		sysconfig.InitEnvVariables("dev")
	}

	// Get Logging information.
	logLevel := sysconfig.GetEnvVariables().GetStringVariable("LogLevel")
	logOutputList := sysconfig.GetEnvVariables().GetListVariable("Logging")
	logOutputListForLoggerError := sysconfig.GetEnvVariables().GetListVariable("LoggingError")
	//Logging=["stdout", "/Volumes/Data/estj/src/estj.log"]
	//LoggingError=["stdout", "/Volumes/Data/estj/src/logger_error.log"]
	// Init logger
	logger.InitLogger(logLevel, logOutputList, logOutputListForLoggerError)

	// Get Database information.
	DBInfo := sysconfig.GetEnvVariables().GetMapVariable("DBInfo")
	if DBInfo == nil {
		createProfileErrors := exception.CreateProfileErrors(reflect.TypeOf(app).String(), "Failed to get database information.")
		log.Fatal(fmt.Sprintf("%+v", errors.Wrap(createProfileErrors, createProfileErrors.GetMessage())))
	}

	// Init Database.
	config.InitDB(DBInfo)

	// Set database.
	dbInstance := config.GetDB()
	if isRunApp {
		defer func(db *sqlx.DB) {
			err := db.Close()
			if err != nil {
				panic(err)
			}
		}(dbInstance)
	}

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
	if isRunApp {
		route.Start()
	}
}
