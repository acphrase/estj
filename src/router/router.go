package router

import (
	"estj/src/exception"
	log "estj/src/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"reflect"
)

// singleton 객체값(pointer)
var routerInstance *gin.Engine

func init() {
	setRouter()
}

func GetRouter() *gin.Engine {
	if routerInstance == nil {
		setRouter()
	}
	return routerInstance
}

func setRouter() {
	routerInstance = gin.Default()
}

func SetRouting(listFunc []func(engine *gin.Engine)) {
	for _, v := range listFunc {
		v(routerInstance)
	}
}

func Start() {
	if routerInstance == nil {
		setRouter()
	}
	err := routerInstance.Run(":9000") // listen and serve on 0.0.0.0:9000 ("localhost:9000")
	if err != nil {
		createInstanceError := exception.CreateInstanceCreationFailed(reflect.TypeOf(routerInstance).String(), "")
		log.Fatal(fmt.Sprintf("%+v", errors.Wrap(createInstanceError, createInstanceError.GetMessage())))
	}
}
