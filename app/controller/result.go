package controller

import (
	"net/http"
	"strconv"

	"SatohAyaka/leaving-match-backend/service"

	"github.com/gin-gonic/gin"
)

func CreateResultHandler(c *gin.Context) {
}

func GetResultHandler(c *gin.Context) {
	busTimePass := c.Param("bustimeId")
	if busTimePass == "" {
		busTimePass = "0"
	}

	busTimeId, err := strconv.ParseInt(busTimePass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	resultService := service.ResultService{}
	results, err := resultService.GetResult(busTimeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get results"})
		return
	}

	c.JSON(http.StatusOK, results)
}
