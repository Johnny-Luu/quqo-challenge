package router

import (
	"github.com/gin-gonic/gin"
	"quqo_challenge/domain/db"
	"quqo_challenge/infrastructure/config"
)

func InitRouter(p *db.Persistence, a config.Configuration) *gin.Engine {
	if a.AppConfig.IsDebugMode() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiR := r.Group("/api")
	ProductRoutes(apiR, p)

	return r
}
