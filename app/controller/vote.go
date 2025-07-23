package controller

import (
	"SatohAyaka/leaving-match-backend/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateVoteHandler(c *gin.Context) {
	busTimePass := c.Param("bustimeId")
	busTimeId, err := strconv.ParseInt(busTimePass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid bustime ID"})
		return
	}

	userIdPass := c.Param("userId")
	userId, err := strconv.ParseInt(userIdPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	previous := false
	nearest := false
	next := false
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
	if err := voteService.CreateVote(busTimeId, userId, previous, nearest, next); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create vote data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "vote created"})
}
