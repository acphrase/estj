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
	userRepository = new(UserRepository)
	userRepository.dbInstance = config.GetDB()
}

func GetUserRepository() *UserRepository {
	return userRepository
}

func (ur *UserRepository) GetAllUser() (*[]entity.User, error) {
	var users []entity.User
	err := ur.dbInstance.Select(&users, "SELECT * FROM public.user")
	return &users, err
}
