package core

import (
	"estj/src/api"
	"estj/src/config"
	"estj/src/router"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func RunApp() {
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
	router.SetRouting(listFunc)
	router.Start()
}
