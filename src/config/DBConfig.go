package config

import (
	"estj/src/core/sysconfig"
	"estj/src/exception"
	log "estj/src/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"reflect"
)

// singleton 객체값(pointer)
var dbInstance *sqlx.DB

func GetDB() *sqlx.DB {
	if dbInstance == nil {
		// Get Database info.
		DBInfo := sysconfig.GetEnvVariables().GetMapVariable("DBInfo")
		if DBInfo == nil {
			createProfileErrors := exception.CreateProfileErrors(reflect.TypeOf(dbInstance).String(), "Failed to get database information.")
			log.Fatal(fmt.Sprintf("%+v", errors.Wrap(createProfileErrors, createProfileErrors.GetMessage())))
		}
		InitDB(DBInfo)
	}
	return dbInstance
}

func InitDB(dbInformation map[string]interface{}) {
	if dbInstance == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbInformation["host"], int(dbInformation["port"].(float64)), dbInformation["user"], dbInformation["password"], dbInformation["dbname"])
		db, err := sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			createDatabaseErrors := exception.CreateDatabaseErrors(reflect.TypeOf(dbInstance).String(), "")
			log.Fatal(fmt.Sprintf("%+v", errors.Wrap(createDatabaseErrors, createDatabaseErrors.GetMessage())))
		}
		//defer func(db *sql.DB) {
		//	err := db.Close()
		//	if err != nil {
		//		panic(err)
		//	}
		//}(dbInstance)
		dbInstance = db
	}
}
