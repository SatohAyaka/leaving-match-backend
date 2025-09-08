package controller

import (
	"SatohAyaka/leaving-match-backend/model"
	"SatohAyaka/leaving-match-backend/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBusTimeHandler(c *gin.Context) {
	recommendedPass := c.Param("recommendedId")
	recommendedId, err := strconv.ParseInt(recommendedPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	previousTime, err := ParseQueryToTime(c.Query("previous"), "previous")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nearestTime, err := ParseQueryToTime(c.Query("nearest"), "nearest")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nextTime, err := ParseQueryToTime(c.Query("next"), "next")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	intPreviousTime, err := strconv.ParseInt(c.Query("previous"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	intEndTime := intPreviousTime - 15
	strEndTime := strconv.FormatInt(intEndTime, 10)
	endTime, err := ParseQueryToTime(strEndTime, "endTime")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bustimeService := service.BusTimeService{}
	lastBusTime, err := bustimeService.GetLatestBusTime()
	if lastBusTime != (model.BusTime{}) {
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid last bustime ID"})
			return
		}
		resultService := service.ResultService{}
		lastResult, err := resultService.GetResult(lastBusTime.BusTimeId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid result ID"})
			return
		}
		if lastResult == (model.Result{}) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "previous bustime does not have result"})
			return
		}
	}

	bustimeId, err := bustimeService.CreateBusTime(recommendedId, previousTime, nearestTime, nextTime, endTime)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create bustime data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"bustime_id": bustimeId})
}

func GetBusTimeByIdHandler(c *gin.Context) {
	busTimeParam := c.Param("bustimeId")
	if busTimeParam == "" {
		busTimeParam = "0"
	}

	busTimeId, err := strconv.ParseInt(busTimeParam, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid bustime ID"})
		return
	}
	busTimeService := service.BusTimeService{}
	bustimes, err := busTimeService.GetBusTime(busTimeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get bustime data"})
		return
	}

	c.JSON(http.StatusOK, bustimes)
}

func GetLatestBusTimeHandler(c *gin.Context) {
	bustimeService := service.BusTimeService{}
	bustime, err := bustimeService.GetLatestBusTime()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get bustime data"})
		return
	}
	c.JSON(http.StatusOK, bustime.BusTimeId)
}

func ParseQueryToTime(query string, errorlabel string) (time.Time, error) {
	if query == "" {
		return time.Time{}, fmt.Errorf("%s bus time choice is required", errorlabel)
	}
	minutes, err := strconv.Atoi(query)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid %s bus time value", errorlabel)
	}
	hour := minutes / 60
	minute := minutes % 60
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location()), nil
}
