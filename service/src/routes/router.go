package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/peya/src/di"
)

var injector *di.Injector

func init() {
	injector = di.InitializeInjector()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", healthCheckHandler)

	return r
}
