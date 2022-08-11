package repository

import (
	"estj/src/config"
	"estj/src/dataaccesslayer/entity"
	"github.com/jmoiron/sqlx"
)

// singleton 객체값(pointer)
var userRepository *UserRepository

type UserRepository struct {
	dbInstance *sqlx.DB
}

func init() {
	initUserRepository()
}

func GetUserRepository() *UserRepository {
	initUserRepository()
	return userRepository
}

func initUserRepository() {
	if userRepository == nil {
		userRepository = new(UserRepository)
	}
}

func (userRepository *UserRepository) getDB() *sqlx.DB {
	if userRepository.dbInstance == nil {
		userRepository.dbInstance = config.GetDB()
	}
	return userRepository.dbInstance
}

func (userRepository *UserRepository) GetAllUser() (*[]entity.User, error) {
	var users []entity.User
	err := userRepository.getDB().Select(&users, "SELECT * FROM public.user")
	return &users, err
}
