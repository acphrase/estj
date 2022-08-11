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
	if userRepository == nil {
		initUserRepository()
	}
	return userRepository
}

func initUserRepository() {
	userRepository = new(UserRepository)
	userRepository.dbInstance = config.GetDB()
}

func (ur *UserRepository) GetAllUser() (*[]entity.User, error) {
	var users []entity.User
	err := ur.dbInstance.Select(&users, "SELECT * FROM public.user")
	return &users, err
}
