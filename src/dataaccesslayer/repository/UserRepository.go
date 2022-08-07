package repository

import (
	"estj/src/config"
	"estj/src/dataaccesslayer/entity"
	"estj/src/exception"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"reflect"
)

// singleton 객체값(pointer)
var instance *UserRepository

type UserRepository struct {
	dbInstance *sqlx.DB
}

func init() {
	instance = new(UserRepository)
	instance.dbInstance = config.GetDB()
}

func GetUserRepository() *UserRepository {
	return instance
}

func (ur *UserRepository) GetAllUser() (*[]entity.User, error) {
	var users []entity.User
	err := ur.dbInstance.Select(&users, "SELECT * FROM public.user")
	if err != nil {
		return nil, errors.Wrap(err, "User is empty")
	}
	if users == nil {
		newError := exception.CreateResourceNotFound(reflect.TypeOf(instance).String(), "")
		fmt.Printf("%+v", errors.Wrap(newError, newError.GetMessage()))
		return nil, newError
	}
	return &users, nil
}
