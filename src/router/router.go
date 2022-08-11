package router

import (
	"estj/src/exception"
	log "estj/src/logger"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"reflect"
	"time"
)

// singleton 객체값(pointer)
var router *Router

type Router struct {
	router *gin.Engine
}

func init() {
	initRouter()
}

func GetRouter() *Router {
	if router == nil {
		initRouter()
	}
	return router
}

func initRouter() {
	router = new(Router)
	router.router = gin.Default()
}

func (router *Router) SetTrustedProxiesPlatforms() {
	err := router.router.SetTrustedProxies([]string{
		"127.0.0.1",
	})
	if err != nil {
		createRouterErrors := exception.CreateRouterErrors(reflect.TypeOf(router.router).String(), "")
		log.Fatal(fmt.Sprintf("%+v", errors.Wrap(createRouterErrors, createRouterErrors.GetMessage())))
	}

	// Use predefined header gin.PlatformXXX
	router.router.TrustedPlatform = gin.PlatformGoogleAppEngine
	// Or set your own trusted request header for another trusted proxy service
	// Don't set it to any suspect request header, it's unsafe
	router.router.TrustedPlatform = "X-CDN-IP"
}

func (router *Router) SetCORS() {
	router.router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"http://localhost:9000"},
			AllowMethods: []string{"POST"},
			AllowHeaders: []string{"Origin", "custom-header"},
			MaxAge:       12 * time.Hour,
		}))
}

func (router *Router) SetRouting(listFunc []func(engine *gin.Engine)) {
	for _, v := range listFunc {
		v(router.router)
	}
}

func (router *Router) Start() {
	if router.router == nil {
		initRouter()
	}
	err := router.router.Run(":9000") // listen and serve on 0.0.0.0:9000 ("localhost:9000")
	if err != nil {
		createInstanceError := exception.CreateInstanceCreationFailed(reflect.TypeOf(router.router).String(), "")
		log.Fatal(fmt.Sprintf("%+v", errors.Wrap(createInstanceError, createInstanceError.GetMessage())))
	}
}
