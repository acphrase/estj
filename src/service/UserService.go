package service

import (
	"estj/src/dataaccesslayer/repository"
	"estj/src/exception"
	log "estj/src/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
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
		repositoryError := errors.Wrap(err, "User repostiroy")
		log.Info(fmt.Sprintf("%+v", repositoryError))
		c.JSON(http.StatusNotFound, gin.H{"error": repositoryError.Error()})
		return
	} else if users == nil {
		userEmptyError := exception.CreateResourceNotFound(reflect.TypeOf(userService).String(), "")
		log.Info(fmt.Sprintf("%+v", errors.Wrap(userEmptyError, userEmptyError.GetMessage())))
		c.JSON(http.StatusNotFound, gin.H{"error": userEmptyError.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
