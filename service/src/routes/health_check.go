package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthCheckHandler(c *gin.Context) {
	res, err := injector.HealthCheckUsecase.GetStatus()
	returnResponseOrFail(c, res, err)
}

func returnResponseOrFail(c *gin.Context, res any, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
