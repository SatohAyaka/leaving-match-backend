package controller

import (
	"SatohAyaka/leaving-match-backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePredictionHandler(c *gin.Context) {
	busTimePass := c.Param("bustimeId")
	busTimeId, err := strconv.ParseInt(busTimePass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid bustime ID"})
		return
	}

	userIdQuery := c.Query("user-id")
	userId, err := strconv.ParseInt(userIdQuery, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	predictionTime, err := ParseQueryToTime(c.Query("time"), "prediction")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	predictionService := service.PredictionService{}
	if err := predictionService.CreatePrediction(busTimeId, userId, predictionTime); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create prediction data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "prediction created"})
}

func GetPredictionHandler(c *gin.Context) {
	busTimePass := c.Param("bustimeId")
	if busTimePass == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid bustime ID"})
		return
	}
	busTimeId, err := strconv.ParseInt(busTimePass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid bustime ID"})
		return
	}

	predictionService := service.PredictionService{}
	predictions, err := predictionService.GetPrediction(busTimeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get prediction Data"})
		return
	}

	c.JSON(http.StatusOK, predictions)
}
