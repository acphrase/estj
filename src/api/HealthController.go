package api

import "github.com/gin-gonic/gin"

func HealthController(engine *gin.Engine) {

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "alive",
		})
	})
}
