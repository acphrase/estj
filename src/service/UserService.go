package service

import (
	"estj/src/dataaccesslayer/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userService *UserService

type UserService struct {
	userRepository *repository.UserRepository
}

func init() {
	userService = new(UserService)
	userService.userRepository = repository.GetUserRepository()
}

func GetUserService() *UserService {
	return userService
}

func (us *UserService) GetAllUser(c *gin.Context) {
	users, err := us.userRepository.GetAllUser()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, users)
}
