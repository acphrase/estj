package api

import (
	"estj/src/service"
	"github.com/gin-gonic/gin"
)

func UserController(engine *gin.Engine) {

	engine.GET("/users", service.GetUserService().GetAllUser)
}
