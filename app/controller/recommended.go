package controller

import (
	"SatohAyaka/leaving-match-backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRecommendedHandler(c *gin.Context) {
	recommendedTime, err := ParseQueryToTime(c.Query("time"), "recommended")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get recommendedTime"})
		return
	}

	membersQuery := c.QueryArray("member")
	memberIds := make([]int, 0, len(membersQuery))
	for _, m := range membersQuery {
		id, err := strconv.Atoi(m)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid member id"})
			return
		}
		memberIds = append(memberIds, id)
	}

	recommendedService := service.RecommendedService{}
	response, err := recommendedService.CreateRecommended(recommendedTime, memberIds)

	c.JSON(http.StatusOK, response)

}

func GetLatestRecommendedStatusHandler(c *gin.Context) {}

func GetLatestRecommendedMembersHandler(c *gin.Context) {}
