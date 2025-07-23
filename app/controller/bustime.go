package controller

import (
	"SatohAyaka/leaving-match-backend/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBusTimeHandler(c *gin.Context) {
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

	members := c.QueryArray("member")
	if len(members) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "at least one member ID is required"})
		return
	}

	memberString := strings.Join(members, ",")

	bustimeService := service.BusTimeService{}
	bustimeId, err := bustimeService.CreateBusTime(memberString, previousTime, nearestTime, nextTime)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create bustime data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"bustime_id": bustimeId})
}

func GetBusTimeHandler(c *gin.Context) {
	busTimeQuery := c.QueryArray("id")

	var busTimeIds []int64
	for _, idStr := range busTimeQuery {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid bustime ID: " + idStr})
			return
		}
		busTimeIds = append(busTimeIds, id)
	}

	busTimeService := service.BusTimeService{}
	bustimeData, err := busTimeService.GetBusTime(busTimeIds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get bustimeData"})
		return
	}

	c.JSON(http.StatusOK, bustimeData)
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
