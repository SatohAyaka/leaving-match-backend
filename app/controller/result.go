package controller

import (
	"net/http"
	"strconv"

	"SatohAyaka/leaving-match-backend/model"
	"SatohAyaka/leaving-match-backend/service"

	"github.com/gin-gonic/gin"
)

func CreateResultHandler(c *gin.Context) {
	busTimePass := c.Param("bustimeId")
	busTimeId, err := strconv.ParseInt(busTimePass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	voteService := service.VoteService{}
	votes, err := voteService.GetVote(busTimeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get votes"})
		return
	}

	previous, nearest, next := VotingResult(votes)

	bustimeService := service.BusTimeService{}
	resultService := service.ResultService{}

	var resultId int64

	if previous >= nearest && previous >= next {
		busTime, err := bustimeService.BusTimeToId(busTimeId, 1)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get bustime"})
			return
		}
		resultId, err = resultService.CreateResult(busTimeId, busTime, previous)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create bustime data"})
			return
		}
	} else if nearest >= next {
		busTime, err := bustimeService.BusTimeToId(busTimeId, 2)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get bustime"})
			return
		}
		resultId, err = resultService.CreateResult(busTimeId, busTime, nearest)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create bustime data"})
			return
		}
	} else {
		busTime, err := bustimeService.BusTimeToId(busTimeId, 3)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get bustime"})
			return
		}
		resultId, err = resultService.CreateResult(busTimeId, busTime, next)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create bustime data"})
			return
		}
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create result"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result_id": resultId})

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

func VotingResult(votes []model.Vote) (previousVote int64, nearestVote int64, nextVote int64) {
	previous := 0
	nearest := 0
	next := 0
	for _, vote := range votes {
		if vote.Previous {
			previous += 1
		}
		if vote.Nearest {
			nearest += 1
		}
		if vote.Next {
			next += 1
		}
	}
	return int64(previous), int64(nearest), int64(next)
}
