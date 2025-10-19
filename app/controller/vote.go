package controller

import (
	"SatohAyaka/leaving-match-backend/model"
	"SatohAyaka/leaving-match-backend/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateVoteHandler(c *gin.Context) {
	busTimeService := service.BusTimeService{}
	lastBusTime, err := busTimeService.GetLatestBusTime()
	busTimeId := lastBusTime.BusTimeId
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid bustime ID"})
		return
	}

	resultService := service.ResultService{}
	lastResult, err := resultService.GetResult(busTimeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid result ID"})
		return
	}
	if lastResult != (model.Result{}) {
		c.JSON(http.StatusForbidden, gin.H{"message": "voting closed for this bustime"})
		return
	}

	slackIdPass := c.Param("slackUserId")
	backendId, err := SlackIdToBackendId(slackIdPass)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	previous, nearest, next := false, false, false
	busTimeQuery := c.Query("vote")
	if busTimeQuery == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "vote query parameter is required"})
		return
	}
	for _, v := range strings.Split(busTimeQuery, ",") {
		switch strings.TrimSpace(v) {
		case "previous":
			previous = true
		case "nearest":
			nearest = true
		case "next":
			next = true
		}
	}

	voteService := service.VoteService{}
	if err := voteService.CreateVote(busTimeId, backendId, previous, nearest, next); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create vote data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "vote created"})
}

func GetVoteHandler(c *gin.Context) {
	busTimePass := c.Param("bustimeId")
	if busTimePass == "" {
		busTimePass = "0"
	}
	busTimeId, err := strconv.ParseInt(busTimePass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid bustime ID"})
		return
	}
	voteService := service.VoteService{}
	votes, err := voteService.GetVote(busTimeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get vote data"})
		return
	}
	c.JSON(http.StatusOK, votes)
}
