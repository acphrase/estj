package test

import (
	"estj/src/config"
	"estj/src/core/sysconfig"
	"estj/src/dataaccesslayer/repository"
	"estj/src/exception"
	log "estj/src/logger"
	"estj/src/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func TestUserService_GetAllUser(t *testing.T) {

	// Given
	// Init environment variables.
	sysconfig.InitEnvVariables("dev")

	// Get Database information.
	DBInfo := sysconfig.GetEnvVariables().GetMapVariable("DBInfo")
	if DBInfo == nil {
		createProfileErrors := exception.CreateProfileErrors("At TestUserService_GetAllUser", "Failed to get database information.")
		log.Fatal(fmt.Sprintf("%+v", errors.Wrap(createProfileErrors, createProfileErrors.GetMessage())))
	}

	// Init Database.
	config.InitDB(DBInfo)

	// Set database.
	dbInstance := config.GetDB()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(dbInstance)

	// When
	ur := repository.GetUserRepository()
	users, err := ur.GetAllUser()

	// Then
	if err != nil {
		repositoryError := errors.Wrap(err, "User repostiroy")
		log.Info(fmt.Sprintf("%+v", repositoryError))
		return
	} else if users == nil {
		us := service.GetUserService()
		userEmptyError := exception.CreateResourceNotFound(reflect.TypeOf(us).String(), "")
		log.Info(fmt.Sprintf("%+v", errors.Wrap(userEmptyError, userEmptyError.GetMessage())))
		return
	}
}
